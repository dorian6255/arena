package handler

import (
	"fmt"
	"github.com/dorian6255/arena/internal/gameLoop"
	"log/slog"
)

type Handler interface {
	Start()
	Stop()
	Init()
}

type BaseHandler struct {
	game    *gameLoop.GameLoop
	outputs []*gameLoop.OuputAdapter
}

// init the handler by allowing adapter & gameLoop
func (b *BaseHandler) InitGameLoop() {
	b.game = &gameLoop.GameLoop{}
	//b.game.Init()
	//TODO

}

func (b *BaseHandler) InitOutputAdapter(adapters ...gameLoop.OuputAdapter) {

	for _, adapter := range adapters {
		err := adapter.Init()
		fmt.Printf("TODO, DO SMTH WITH ERROR %v ", err)
		b.outputs = append(b.outputs, &adapter)
	}
	slog.Info("OutputAdapter Initated")
}
