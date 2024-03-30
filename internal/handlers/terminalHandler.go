package handler

import (
	"log/slog"
)

type TerminalHandler struct {
	BaseHandler
}

func (t TerminalHandler) Init() {
	//init
	t.InitGameLoop()
	//t.InitOutputAdapter(myadapter)
	slog.Info("Terminal Handler Initated")
}

func (t TerminalHandler) Start() {
	//init
	slog.Info("Terminal Handler Started")
	//display welcom message and infos
	//loop and wait for reply
	// process
	// show result
	// show leaderboard  in a if?
}

func (t TerminalHandler) Stop() {

	slog.Info("Terminal Handler Stop")

}
