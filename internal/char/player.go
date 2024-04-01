package char

import (
	"errors"
)

// implementation of Char to use a the Player
type Player struct {
	Char
}

// assign stats and calcul HP
// this one is for the Player and not the other, because we want to be able to create enemies that doesn't follow the rules for stats
// it call Char.Init for stat assignation
func (p *Player) InitPlayer(str, dex, con, inte, wis, cha, lvl int) error {
	if (str + dex + con + inte + wis + cha) != 27 {
		return errors.New("Total stats points is diffrent than 27")
	}
	values := [6]int{str, dex, con, inte, wis, cha}

	countNine := 0
	countEight := 0
	for _, value := range values {
		if value > 9 {
			return errors.New("Value Superior to 9")
		} else if value < 0 {
			return errors.New("cannot have value under 0")
		}
		switch value {
		case 9:
			if countNine == 1 {
				return errors.New("Too many 9")
			} else {
				countNine++
			}
		case 8:
			if countEight == 1 {
				return errors.New("Too many 8")
			} else {
				countEight++
			}
		}

	}

	p.Init(str, dex, con, inte, wis, cha, lvl)
	p.baseDmg = 2

	return nil
}
