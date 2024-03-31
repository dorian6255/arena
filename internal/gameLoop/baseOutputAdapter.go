package gameLoop

import ()

type OuputAdapter interface {
	Init() error
	SendStartingMessage() error
	SendResultMessage(game *GameLoop) error
	SendFinishMessage() error
	SendLeaderBoard() error
}
