package char

import "fmt"

const MAX_ALLO_STATS = 27 //Based on baldur's gate 3 character creator

// stats of the player and also the enemies
type stats struct {
	str  int
	dex  int
	con  int
	inte int
	wis  int
	cha  int
}

// stats + HP, and methods
type Char struct {
	stats
	hp int
}

// implementation of Char to use a the player
type Player struct {
	Char
}

// assign stats and calcul HP
// this one is for the player and not the other, because we want to be able to create enemies that doesn't follow the rules for stats
// it call Char.Init for stat assignation
func (p *Player) InitPlayer(str, dex, con, inte, wis, cha int) error {
	return nil
}

// this one is useful to create enemies without checking if the stats are valid
func (p *Char) Init(str, dex, con, inte, wis, cha int) error {
	return nil
}

// verify that the player is valid (== stats are ok)
func (p *Char) isValid() bool {

	return false
}

// we need a way to check if HP > 0
func (p *Char) IsAlive() bool {
	return false
}

// calcul if hit then trigger target's ReceiveDamage methods with dmg
func (p *Char) Attacks(target *Char) error {
	fmt.Println(target)
	return nil
}

// ReceiveDamage
func (p *Char) ReceiveDamage(dmg int) error {
	fmt.Println(dmg)
	return nil
}
