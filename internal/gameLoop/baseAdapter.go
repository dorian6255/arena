package gameLoop

import ()

// interface of outputadapter, that allow to send message to multiple and different kind of output interface
// discord, messenge, terminal, ...
type OuputAdapter interface {
	sendStartingMessage() error
	sendResultMessage(game *GameLoop) error
	sendFinishMessage() error
	sendLeaderBoard(game *GameLoop) error
}

