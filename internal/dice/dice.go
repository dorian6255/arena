package dice

import (
	"math/rand"
)

type Dice struct {
	max int //numberMax on dice //d20 = 20, ...
}

func (d *Dice) Roll() (res int) {
	if d.max <= 0 {
		return 0
	}
	return rand.Intn(d.max)
}
