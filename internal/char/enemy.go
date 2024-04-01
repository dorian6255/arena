package char

import (
	"github.com/dorian6255/arena/internal/dice"
	"log/slog"
)

type Enemy struct {
	Char
}

// The base enemy is a char with low stat, we roll a 8 dices and we remove the result for each stat
// we'll add specials enemy later
func (e *Enemy) InitBaseEnemy() error {
	dice := dice.Dice{Max: 2}
	//8
	// 4 2
	e.Init(-dice.Roll()-dice.Roll(), -dice.Roll()-dice.Roll(), -dice.Roll()-dice.Roll(), -dice.Roll()-dice.Roll(), -dice.Roll()-dice.Roll(), -dice.Roll()-dice.Roll(), dice.Roll()+2)
	e.baseDmg = 1
	slog.Debug("Initated enemy with ", "stats", e.getStats())
	return nil
}
