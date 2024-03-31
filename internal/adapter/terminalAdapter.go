package adapter

import (
	"fmt"
	"github.com/dorian6255/arena/internal/gameLoop"
)

// Adapter to play in terminal and see the result
type TerminalAdapter struct {
}

func (d TerminalAdapter) Init() error {
	//Nothing to do here
	return nil
}
func (d TerminalAdapter) SendStartingMessage() error {

	fmt.Println("Welcome to the Arena")

	return nil
}

func (d TerminalAdapter) SendFinishMessage() error {
	fmt.Println("Finished !")
	return nil
}

func (d TerminalAdapter) SendResultMessage(game *gameLoop.GameLoop) error {
	//TODO
	fmt.Println(game)

	return nil
}

func (d TerminalAdapter) SendLeaderBoard() error {
	return nil
}
