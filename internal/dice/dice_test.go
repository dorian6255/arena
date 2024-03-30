package dice

import (
	"testing"
)

func TestRollValueInInterval(t *testing.T) {

	var tests = []struct {
		name  string
		input int
	}{
		{"Roll with 10", 10},
		{"Roll with 8", 8},
		{"Roll with 100", 100},
		{"Roll with 20", 20},
		{"Roll with 30", 30},
		{"Roll with 0", 0},
		{"Roll with -10", -10},
		{"Roll with -100", -100},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			r := Dice{tt.input}

			for i := 0; i < 100; i++ { //let's try 100 hundred time for each, to be sure...
				res := r.Roll()
				if tt.input > 0 {
					if !(res >= 0 && res <= tt.input) {
						t.Errorf("dice rolled value out of range %v, dice was %v ", res, tt.input)
					}

				} else {
					if res != 0 {
						t.Errorf("dice rolled %v, but need to roll 0 because dice was %v ", res, tt.input)
					}
				}
			}
		})
	}

	t.Run("TestRollDifferentValues", func(t *testing.T) {
		r := Dice{20}
		allRes := map[int]int{}
		for i := 0; i < 200; i++ { //let's try 200 hundred time for each, to be sure...
			res := r.Roll()
			allRes[res] = allRes[res] + 1
		}

		for i := 0; i < 20; i++ {
			if allRes[i] == 0 {
				t.Errorf("dice never rolled %v ", i)
			}
		}
	})

}
