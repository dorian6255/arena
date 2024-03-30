package handler

import (
	"log/slog"
)

type TerminalHandler struct {
	BaseHandler
}

func (t *TerminalHandler) Init() {
	//init
	t.InitGameLoop()
	//t.InitOutputAdapter(myadapter)
	slog.Info("Terminal Handler Initated")
}

func (t *TerminalHandler) Start() {
	//init
	slog.Info("Terminal Handler Started")
	//display welcom message and infos

	//TODO ask player for info about PLAYER
	t.game.Init(1, 2, 3, 8, 7, 6, 50)
	//loop and wait for reply
	t.game.Process()
	// process
	// show result
	// show leaderboard  in a if?
}

func (t TerminalHandler) Stop() {

	slog.Info("Terminal Handler Stop")

}
