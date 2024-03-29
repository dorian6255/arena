package adapter

import (
	"fmt"
	"github.com/dorian6255/arena/internal/gameLoop"
)

// TODO
type DiscordAdapter struct {
}

func (d *DiscordAdapter) sendStartingMessage() error {
	//TODO
	return nil
}
func (d *DiscordAdapter) sendFinishMessage() error {
	//TODO
	return nil
}

func (d *DiscordAdapter) sendResultMessage(game *gameLoop.GameLoop) error {
	//TODO
	fmt.Println(game)
	return nil
}
func (d *DiscordAdapter) sendLeaderBoard(game *gameLoop.GameLoop) error {

	fmt.Println(game)
	return nil
}
