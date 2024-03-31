package gameLoop

import (
	"github.com/dorian6255/arena/internal/char"

	"log/slog"
)

// each round is stored as a struct to allow a detailled log of every game
type round struct {
	enemies []char.Char
}

func (g *GameLoop) createRound(idxRound int) {

}

// all the logic and data of the game is contained in this superStruct Gameloop
type GameLoop struct {
	goal   int //not mandatory
	player *char.Player
	rounds []round
}

// verify that the Player valid, then add it as the player
func (g *GameLoop) initPlayer(str, dex, con, inte, wis, cha int) error {

	p := char.Player{}
	err := p.InitPlayer(str, dex, con, inte, wis, cha)
	if err != nil {
		return err
	}
	g.player = &p

	slog.Info("Player Initated")
	return nil
}

// main loop
func (g *GameLoop) Process() error {
	//startingRound := [6]char.Char{}
	for round := 0; round != g.goal && g.player.IsAlive(); round++ {
		//TODO create round (filled enemies )
		slog.Info("Round ", "round", round)
		//TODO fight until everyone on one side is dead
	}

	return nil
}

// init players, enemies, nbround, adapter, ... and check that everything is alright before starting
func (g *GameLoop) Init(str, dex, con, inte, wis, cha, goals int) error {
	slog.Info("Starting to Init GameLoop")
	err := g.initPlayer(str, dex, con, inte, wis, cha)
	if err == nil {

		g.goal = goals
		slog.Info("GameLoop Initated")
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
	return g.player.IsAlive()
}
