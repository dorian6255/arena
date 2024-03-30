package gameLoop

import (
	"github.com/dorian6255/arena/internal/char"
)

// each round is stored as a struct to allow a detailled log of every game
type round struct {
	enemies []char.Char
}

// all the logic and data of the game is contained in this superStruct Gameloop
type GameLoop struct {
	goal          int //not mandatory
	player        *char.Player
	outputAdapter *OuputAdapter
	inputAdapter  *InputAdapter
	rounds        []round
}

// verify that the Player valid, then add it as the player
func (g *GameLoop) initPlayer(player *char.Player) error {
	g.player = player

	return nil
}

// main loop
func (g *GameLoop) Process() error {
	return nil
}

// init players, enemies, nbround, adapter, ... and check that everything is alright before starting
func (g *GameLoop) Init() error {

	return nil
}

// go to next round
func (g *GameLoop) newRound() {

}

// check if player is dead or if goal is met
func (g *GameLoop) isFinished() bool {
	return g.player.IsAlive()
}
