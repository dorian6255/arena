package char

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInit(t *testing.T) {

	p := Player{}
	p.Init(0, 0, 0, 0, 0, 0)

	//default value
	if !reflect.DeepEqual(p.getStats(), [6]int{8, 8, 8, 8, 8, 8}) {
		t.Errorf("Default stats values should be {8 8 8 8 8 8} got %v", p.stats)
	}
	//default HP
	if p.hp != 7 { // 8 - 1 because con modifier is -1 for con = 1
		t.Errorf("Base HP should Be 7")
	}
}

func TestInitPlayerWrongValues(t *testing.T) {
	var tests = []struct {
		name  string
		input [6]int
	}{
		{"should not accept negative value", [6]int{-1, -1, -1, -1, -1, -1}},
		{"should not accept sum value of more than " + strconv.Itoa(MAX_ALLO_STATS), [6]int{(MAX_ALLO_STATS / 6) + 1, (MAX_ALLO_STATS / 6) + 1, (MAX_ALLO_STATS / 6) + 1, (MAX_ALLO_STATS / 6) + 1, (MAX_ALLO_STATS / 6) + 1, MAX_ALLO_STATS/6 + 1}},
		{"should not accept values superior to 9 (because 17 is the max you can have) ", [6]int{0, 0, 0, 0, 0, 27}},
		{"should not accept values superior to 9 (because 17 is the max you can have) ", [6]int{3, 3, 6, 2, 10, 3}},
		{"every point should be spend", [6]int{0, 0, 0, 0, 0, 0}},
		{"only one 9 is possible", [6]int{9, 9, 3, 3, 3, 0}},
		{"only one 8 is possible", [6]int{8, 8, 4, 3, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := Player{}
			err := p.InitPlayer(tt.input[0], tt.input[1], tt.input[2], tt.input[3], tt.input[4], tt.input[5])
			if err == nil {
				t.Errorf("InitPlayer(%v) should have returned an error, but didn't err:%v ", tt.input, err)
			}

		})
	}
}

func TestGetModifier(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  int
	}{
		{"modifier with 10", 10, 0},
		{"modifier with 8", 8, -1},
		{"modifier with 4", 4, -3}, // for debuff
		{"modifier with 15", 15, 2},
		{"modifier with 17", 17, 3},
		{"modifier with 20", 20, 5},
		{"modifier with 26", 26, 8}, // for buff
		{"modifier with 30", 30, 10},
		{"modifier with 31", 31, 10},
		{"modifier with 100", 100, 10}, // no more than +10
		{"modifier with 1", 1, -5},
		{"modifier with -1", -1, -5},     // no less than -5
		{"modifier with -100", -100, -5}, // no less than -5
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := GetModifier(tt.input)
			if tt.want != got {
				t.Errorf("GetModifier(%v) got invalid modifier : %v instead of %v ", tt.input, got, tt.want)
			}
		})
	}
}
func TestInitPlayer(t *testing.T) {
	var tests = []struct {
		name  string
		input [6]int
		want  [6]int
	}{
		{"init with 1 2 3 8 7 6", [6]int{1, 2, 3, 8, 7, 6}, [6]int{9, 10, 11, 16, 15, 14}},
		{"init with a 0 ", [6]int{0, 3, 3, 8, 7, 6}, [6]int{8, 11, 11, 16, 15, 14}},
		{"init with a 9 ", [6]int{0, 3, 2, 9, 7, 6}, [6]int{8, 11, 10, 17, 15, 14}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Player{}
			err := p.InitPlayer(tt.input[0], tt.input[1], tt.input[2], tt.input[3], tt.input[4], tt.input[5])

			if err != nil {
				t.Errorf("InitPlayer(%v) returned an error, but but shouldn't have: %v", tt.input, err)
			}
			got := p.getStats()
			if !reflect.DeepEqual(got, tt.want) {

				t.Errorf("InitPlayer(%v) set wrong stats values, want: %v, got : %v", tt.input, tt.input, got)
			}

		})

	}

}

func TestReceiveDamage(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  int
		error bool
	}{ //given  player have 7hp
		{"ReceiveDamage 7", 7, 0, false},
		{"ReceiveDamage -1", -1, 7, true},
		{"ReceiveDamage 0", 0, 7, false},
		{"ReceiveDamage 100", 100, -93, false},
	}

	for _, tt := range tests {
		p := Player{}
		p.Init(0, 0, 0, 0, 0, 0) // 7 hp

		t.Run(tt.name, func(t *testing.T) {
			err := p.receiveDamage(tt.input)
			if tt.error && err == nil {
				t.Errorf("ReceiveDamage(%v) did not give error but should have ", tt.input)
			}
			if p.hp != tt.want {
				t.Errorf("ReceiveDamage(%v) did not remove th execpted (%v) amount of HP, hp : %v", tt.input, tt.want, tt.input)
			}
		})

	}
}

func TestIsAlive(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  bool
	}{
		{"IsAlive 7 hp", 0, true},
		{"IsAlive 0 hp ", 7, false},
		{"IsAlive -1 hp ", 8, false},
	}

	for _, tt := range tests {
		p := Player{}
		p.Init(0, 0, 0, 0, 0, 0) // 7 hp

		p.receiveDamage(tt.input)
		if tt.want != p.IsAlive() {
			t.Errorf("IsAlive give wrong result for dmg hp %v, gave : %v, hp: %v ", tt.input, p.IsAlive(), p.hp)
		}

	}

}
