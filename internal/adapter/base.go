package adapter

import (
	"github.com/dorian6255/arena/char"
)

type Adapter interface {
	createPlayer() char.Player
}
