package adapter

import (
	"github.com/dorian6255/arena/internal/char"
)

type Adapter interface {
	createPlayer() char.Player
}
