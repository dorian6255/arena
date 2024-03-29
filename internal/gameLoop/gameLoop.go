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
	round         []round
}
