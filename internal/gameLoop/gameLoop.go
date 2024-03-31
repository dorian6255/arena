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

func findBestInitiative(challengers map[*char.Char]int) (winner *char.Char) {
	bestInitiative := -1000

	for challenger, initiative := range challengers {
		if initiative > bestInitiative {
			bestInitiative = initiative
			winner = challenger
		}
	}
	return
}

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

func (r *Round) resolveRound(player *char.Player, playerInitiative int) (nbTurn int) {
	for player.IsAlive() && !r.roundDefeated() {
		r.newTurn(player, playerInitiative)
		nbTurn++
	}
	return
}

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
func (g *GameLoop) createRound(idxRound int) {

	newRound := Round{}
	for i := 0; i < idxRound; i++ {
		newEnemy := char.Enemy{}
		newEnemy.InitBaseEnemy()
		newRound.enemies = append(newRound.enemies, &newEnemy.Char)
	}
	g.rounds = append(g.rounds, newRound)

	newRound.setFightOrderEnemies()

}

// all the logic and data of the game is contained in this superStruct Gameloop
type GameLoop struct {
	goal   int //not mandatory
	player *char.Player
	rounds []Round
}

// verify that the Player valid, then add it as the player
func (g *GameLoop) initPlayer(str, dex, con, inte, wis, cha, lvl int) error {

	p := char.Player{}
	err := p.InitPlayer(str, dex, con, inte, wis, cha, lvl)
	if err != nil {
		return err
	}
	g.player = &p

	slog.Debug("Player Initated")
	return nil
}

// main loop
func (g *GameLoop) Process() error {

	for i := 0; i != g.goal && g.player.IsAlive(); i++ {
		//create round (filled enemies )
		g.player.Rest()
		slog.Info("Round ", "round", i+1)
		g.createRound(i + 1)
		slog.Debug("NbEnemies ", "nb", len(g.rounds[i].enemies))
		pInitiative := g.player.RollInitiative(0)
		slog.Debug("Player initiative roll ", "pInitiative", pInitiative)
		nbTurn := g.rounds[i].resolveRound(g.player, pInitiative)

		slog.Info("turn to complete", "nbturn", nbTurn)
		//fight until everyone on one side is dead
		if g.player.IsAlive() {

			slog.Info("Player won round !!")
		} else {
			slog.Info("Player is dead !!")
		}

		slog.Info("nbRounds ", "nbRounds", len(g.rounds))
		resume := g.player.Char.GetResume()

		slog.Info("resume ", "ca, lvl, gothit, hit, missed, crit, dmgDone, dmgTaken", resume)

	}

	return nil
}

// init players, enemies, nbround, adapter, ... and check that everything is alright before starting
func (g *GameLoop) Init(str, dex, con, inte, wis, cha, goals int) error {
	slog.Debug("Starting to Init GameLoop")
	err := g.initPlayer(str, dex, con, inte, wis, cha, 10)
	if err == nil {

		g.goal = goals
		slog.Debug("GameLoop Initated")
		return nil

	} else {
		return err
	}

}

// go to next round
func (g *GameLoop) newRound() {

}

// check if player is dead or if goal is met
func (g *GameLoop) isFinished() bool {
	slog.Debug("isFinish: ", "isFinish", g.player.IsAlive)
	return g.player.IsAlive()
}
