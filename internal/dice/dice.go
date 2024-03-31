package dice

import (
	"log/slog"
	"math/rand"
)

type Dice struct {
	Max int //numberMax on dice //d20 = 20, ...
}

func (d *Dice) Roll() (res int) {
	if d.Max <= 0 {
		return 0
	} else if d.Max == 1 {
		return 1
	}

	maxi := d.Max - 1

	res = rand.Intn(maxi)
	res += 1

	slog.Debug("rolled an ", "dice roll", res)
	return
}
