package char

import (
	"errors"
	"github.com/dorian6255/arena/internal/dice"
	"log/slog"
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
	hp       int
	ca       int
	baseDmg  int
	Name     string
	lvl      int
	gotHit   int
	hit      int
	missed   int
	crit     int
	dmgDone  int
	dmgTaken int
}

func (c *Char) GetResume() (resume []int) {
	//TODO find a way to simplify this
	resume = append(resume, c.ca)
	resume = append(resume, c.lvl)
	resume = append(resume, c.gotHit)
	resume = append(resume, c.hit)
	resume = append(resume, c.missed)
	resume = append(resume, c.crit)
	resume = append(resume, c.dmgDone)
	resume = append(resume, c.dmgTaken)

	return
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

func (c *Char) RollInitiative(malus int) int {
	dice := dice.Dice{Max: 20}

	return dice.Roll() - malus + GetModifier(c.dex)
}

func (c *Char) Rest() {
	c.hp = 8 + (GetModifier(c.con)*c.lvl + c.lvl*6)
	slog.Debug("rested ", "hp", c.hp, "ca", c.ca)
}

// this one is useful to create enemies without checking if the stats are valid
func (c *Char) Init(str, dex, con, inte, wis, cha, lvl int) {
	c.str = str + 8
	c.dex = dex + 8
	c.con = con + 8
	c.inte = inte + 8
	c.wis = wis + 8
	c.cha = cha + 8
	c.lvl = lvl
	c.hp = 8 + (GetModifier(c.con)*lvl + lvl*6)
	slog.Debug("char has ", "hp", c.hp)
	c.ca = GetModifier(c.con) + GetModifier(c.dex) + 10

	slog.Debug("char has ", "ca", c.ca)
}

// we need a way to check if HP > 0
func (c *Char) IsAlive() bool {

	return c.hp > 0
}

// calcul if hit then trigger target's ReceiveDamage methods with dmg
func (p *Char) Attacks(target *Char) error {
	//if it call receivedmg on target with dmg
	//TODO change
	hitDice := dice.Dice{Max: 20}

	rolledHit := hitDice.Roll()
	rolledHit += GetModifier(p.str)
	if rolledHit < 0 {
		rolledHit = 0
	}
	if rolledHit > target.ca {
		slog.Debug("HIT")
		p.hit += 1
		//TODO change calcul
		dmgModifier := (p.str + (GetModifier(p.str) * p.str))
		if dmgModifier < 0 {
			dmgModifier = 0
		}
		dommage := p.baseDmg + dmgModifier
		rolledDamage := dommage
		if rolledHit >= 20-GetModifier(p.wis) {

			p.crit += 1
			//rolledDamage := damageDice.Roll()
			slog.Debug("CRITIQUAL HIT", "rolledDamage", rolledDamage)
			target.receiveDamage(rolledDamage)
			p.dmgDone += rolledDamage
		} else {
			//rolledDamage := damageDice.Roll()
			slog.Debug("Normal Hit done ", "rolledDamage", rolledDamage)
			target.receiveDamage(rolledDamage)
			p.dmgDone += rolledDamage

		}
	} else {
		p.missed += 1
		slog.Debug("missed ", "rolledHit", rolledHit)
	}

	return nil
}

// ReceiveDamage
func (c *Char) receiveDamage(dmg int) error {
	//TODO take armor into consideration via another function
	if dmg >= 0 {
		slog.Debug("receive dmg ", "dmg", dmg)
		slog.Debug("target hp ", "hp", c.hp)
		c.hp -= dmg
		c.dmgTaken += dmg
		c.gotHit += 1
		return nil
	}
	return errors.New("Cannot deal negative amount of damage")
}
