package gameLoop

import (
	"github.com/dorian6255/arena/internal/char"

	"log/slog"
)

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

		slog.Info("Round ", "round", i+1)
		g.newRound(i)
		//fight until everyone on one side is dead
		if g.player.IsAlive() {

			slog.Info("Player won round !!")
		} else {
			slog.Info("Player is dead !!")
		}

	}

	slog.Info("nbRounds ", "nbRounds", len(g.rounds))
	resume := g.player.Char.GetResume()

	slog.Info("resume ", "ca, lvl, gothit, hit, missed, crit, dmgDone, dmgTaken", resume)

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
func (g *GameLoop) newRound(idxRound int) {
	//create round (filled enemies )
	g.player.Rest()
	g.createRound(idxRound + 1)
	slog.Debug("NbEnemies ", "nb", len(g.rounds[idxRound].enemies))
	pInitiative := g.player.RollInitiative(0)
	slog.Debug("Player initiative roll ", "pInitiative", pInitiative)
	g.rounds[idxRound].resolveRound(g.player, pInitiative)

}

// check if player is dead or if goal is met
func (g *GameLoop) isFinished() bool {
	slog.Debug("isFinish: ", "isFinish", g.player.IsAlive)
	return g.player.IsAlive()
}
