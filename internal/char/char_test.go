package char

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestInit(t *testing.T) {

	p := Player{}
	//default value
	if !reflect.DeepEqual(p.stats, [6]int{8, 8, 8, 8, 8, 8}) {
		t.Errorf("Default stats values should be {8 8 8 8 8 8} got %v", p.stats)
	}
	//default HP
	if p.hp != 8 {
		t.Errorf("Base HP should Be 8")
	}
	fmt.Println(p.stats)
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
			if err != nil {
				t.Errorf("InitPlayer(%v) should have returned an error, but didn't err:%v ", tt.input, err)
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
