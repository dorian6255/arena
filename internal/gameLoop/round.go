package gameLoop

import (
	"github.com/dorian6255/arena/internal/char"
	"github.com/dorian6255/arena/internal/dice"
	"log/slog"
)

// each round is stored as a struct to allow a detailled log of every game
type Round struct {
	enemies        []*char.Char
	fightOrder     []*char.Char
	initiativeRoll map[*char.Char]int
}

// set the round.initiativeRoll by rolling dices for each enemies
func (r *Round) setFightOrderEnemies() {
	//initiative = 1 d 20 + dex modifier
	initiativeRoll := make(map[*char.Char]int, len(r.enemies))

	for _, enemy := range r.enemies {
		initiativeRoll[enemy] = enemy.RollInitiative(0)
	}
	r.initiativeRoll = initiativeRoll

	for i := 0; i < len(initiativeRoll); i++ {
		next := findBestInitiative(initiativeRoll)
		r.fightOrder = append(r.fightOrder, next)
		delete(initiativeRoll, next)
	}

}

// return a list of the remaining aliveEnemies in r.enemies
func (r *Round) getAliveFighters() (aliveFighters []*char.Char) {

	for _, fighter := range r.enemies {
		if fighter.IsAlive() {
			aliveFighters = append(aliveFighters, fighter)
		}
	}
	slog.Debug("toFight", "nb", len(aliveFighters))
	return
}

func (r *Round) newTurn(player *char.Player, playerInitiative int) {

	slog.Debug("New turn")
	toFight := r.getAliveFighters()
	slog.Debug("Enemies alives", "toFight", len(toFight))
	didNotPlayAlready := make([]*char.Char, len(r.enemies))
	didNotPlayAlready = r.enemies
	playerPlayed := false

	for len(didNotPlayAlready) != 0 {
		//check if next is alive
		next := didNotPlayAlready[0]
		if next.IsAlive() {
			if !playerPlayed && playerInitiative >= r.initiativeRoll[next] {
				slog.Debug("Player Attacks")
				playerPlayed = true
				max := len(toFight) - 1
				slog.Debug("choose enemies target", "len enemies", len(toFight))
				dice := dice.Dice{Max: max}
				rolledDice := dice.Roll()
				target := toFight[rolledDice]
				player.Attacks(target)
			} else {
				slog.Debug("Enemy Attacks")
				next.Attacks(&player.Char)
				slog.Debug("player HP", "HP", player.Char)
				if len(didNotPlayAlready) != 1 {
					didNotPlayAlready = didNotPlayAlready[1:]
				} else {
					didNotPlayAlready = []*char.Char{}
				}
			}

		} else {
			if len(didNotPlayAlready) != 1 {
				didNotPlayAlready = didNotPlayAlready[1:]
			} else {
				didNotPlayAlready = []*char.Char{}
			}

		}
	}
}

// make the player and enemies fight until one side is dead
func (r *Round) resolveRound(player *char.Player, playerInitiative int) (nbTurn int) {
	for player.IsAlive() && !r.roundDefeated() {
		r.newTurn(player, playerInitiative)
		//put a limit to avoid infinite loop ?
		nbTurn++
	}
	return
}

// check if their are enemies left alive
func (r *Round) roundDefeated() bool {
	for _, enemy := range r.enemies {
		if enemy.IsAlive() {

			slog.Debug("Round not defeated")
			return false
		}
	}
	slog.Debug("Round defeated !!")
	return true
}
