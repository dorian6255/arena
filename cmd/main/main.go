package main

import (
	"github.com/dorian6255/arena/internal/handlers"
)

func main() {

	//TODO getHandlertype from launch command
	//for now we'll use terminal type handler only
	terminalHandler := &handler.TerminalHandler{}

	handlers := []handler.Handler{}
	handlers = append(handlers, terminalHandler)
	//TODO launch each one in a process Goroutine ?
	for _, handler := range handlers {

		handler.Init()
		handler.Start()

	}

	// loop waiting for ctrl-C
	//then close

	handlers[0].Stop()
}
