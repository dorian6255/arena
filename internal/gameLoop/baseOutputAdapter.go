package gameLoop

import ()

type OuputAdapter interface {
	sendStartingMessage() error
	sendResultMessage(game *GameLoop) error
	sendFinishMessage() error
}

