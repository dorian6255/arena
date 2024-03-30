package char

import (
	"errors"
	"fmt"
	"math"
)

const MAX_ALLO_STATS = 27 //Based on baldur's gate 3 character creator

// stats of the Player and also the enemies
type stats struct {
	str  int
	dex  int
	con  int
	inte int
	wis  int
	cha  int
}

func (s *stats) getStats() (res [6]int) {

	res = [6]int{s.str, s.dex, s.con, s.inte, s.wis, s.cha}
	return

}

// stats + HP, and methods
type Char struct {
	stats
	hp int
	ca int
}

// implementation of Char to use a the Player
type Player struct {
	Char
}

// assign stats and calcul HP
// this one is for the Player and not the other, because we want to be able to create enemies that doesn't follow the rules for stats
// it call Char.Init for stat assignation
func (p *Player) InitPlayer(str, dex, con, inte, wis, cha int) error {
	if (str + dex + con + inte + wis + cha) != 27 {
		return errors.New("Total stats points is diffrent than 27")
	}
	values := [6]int{str, dex, con, inte, wis, cha}

	countNine := 0
	countEight := 0
	for _, value := range values {
		if value > 9 {
			return errors.New("Value Superior to 9")
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

	err := p.Init(str, dex, con, inte, wis, cha)
	return err
}

func GetModifier(stat int) (res int) {

	res = int(math.Floor(float64((stat - 10)) / 2))

	// 1 - 10 == -9, -9/2 == -4,5 => 5
	if res < -5 {
		res = -5
	} else if res > 10 {
		res = 10
	}
	return
}

// this one is useful to create enemies without checking if the stats are valid
func (p *Char) Init(str, dex, con, inte, wis, cha int) error {
	if str >= 0 && dex >= 0 && con >= 0 && inte >= 0 && wis >= 0 && cha >= 0 {

		p.str = str + 8
		p.dex = dex + 8
		p.con = con + 8
		p.inte = inte + 8
		p.wis = wis + 8
		p.cha = cha + 8
		p.hp = 8 + GetModifier(p.con)
		return nil
	}
	return errors.New("One or more value is negative")
}

// we need a way to check if HP > 0
func (p *Char) IsAlive() bool {

	return p.hp > 0
}

// calcul if hit then trigger target's ReceiveDamage methods with dmg
func (p *Char) Attacks(target *Char) error {
	//TODO rolldice
	//if it call receivedmg on target with dmg
	fmt.Println(target)
	return nil
}

// ReceiveDamage
func (p *Char) receiveDamage(dmg int) error {
	//TODO take armor into consideration via another function
	if dmg >= 0 {

		p.hp -= dmg
		return nil
	}
	return errors.New("Cannot deal negative amount of damage")
}
