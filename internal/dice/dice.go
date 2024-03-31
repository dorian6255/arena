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
	}
	res = rand.Intn(d.Max)

	slog.Debug("rolled an ", "dice roll", res)
	return
}
