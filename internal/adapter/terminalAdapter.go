package adapter

import (
	"fmt"
	"github.com/dorian6255/arena/internal/gameLoop"
)

// Adapter to play in terminal and see the result
type TerminalAdapter struct {
}

func (d *TerminalAdapter) sendStartingMessage() error {

	fmt.Println("Welcome to the Arena")

	return nil
}

func (d *TerminalAdapter) sendFinishMessage() error {
	fmt.Println("Finished !")
	return nil
}

func (d *TerminalAdapter) sendResultMessage(game *gameLoop.GameLoop) error {
	//TODO
	fmt.Println(game)

	return nil
}

func (d *TerminalAdapter) sendLeaderBoard(game *gameLoop.GameLoop) error {
	fmt.Println(game)
	return nil
}
