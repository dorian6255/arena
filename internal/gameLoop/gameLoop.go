package gameLoop

import (
	"github.com/dorian6255/arena/internal/adapter"
	"github.com/dorian6255/arena/internal/char"
)

type gameLoop struct {
	round         int
	goal          int //not mandatory
	player        *char.Player
	enemies       []char.Char
	outputAdapter *adapter.Adapter
}
