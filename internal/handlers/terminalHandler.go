package handler

import (
	"fmt"
	"github.com/dorian6255/arena/internal/adapter"
	"github.com/dorian6255/arena/internal/char"
	"log/slog"
)

type TerminalHandler struct {
	BaseHandler
}

func (t *TerminalHandler) Init() {
	//init
	t.InitGameLoop()
	// for now just terminal output
	touput := adapter.TerminalAdapter{}
	t.outputs = append(t.outputs, touput)
	//t.InitOutputAdapter(myadapter)
	slog.Info("Terminal Handler Initated")
}

func (t *TerminalHandler) Start() {
	//init
	slog.Info("Terminal Handler Started")
	t.SendStartingMessageToAllAdapters()
	//display welcom message and infos

	//TODO ask player for info about PLAYER
	initated := false
	for !initated {

		err := t.game.Init(getPlayerStats())
		if err == nil {
			fmt.Printf("game initated   \n ")
			initated = true

		} else {
			fmt.Printf("\n\n\nERROR ENCOUNTERED DURING CHARACTER CREATION :  %v \n\n\nRESTARTING \n\n\n", err)
		}
	}
	//loop and wait for reply
	t.game.Process()
	// process
	// show result
	t.SendFinishMessageToAllAdapters()
	t.SendLeaderBoardToAllAdapter()
	// show leaderboard  in a if?
}

func (t TerminalHandler) Stop() {

	slog.Info("Terminal Handler Stop")

}

func getPlayerStats() (str, dex, con, inte, wis, cha, goal int) {

	availablePoints := char.MAX_ALLO_STATS
	fmt.Println("Your character starts with 8 points in every of those stastistics :")
	fmt.Println("STRENGTH DEXTERITY CONSTITUTION INTELLIGENCE WISDOM AND CHARISMA")
	fmt.Printf("You will now be asked to allocate %v points to your statistics", availablePoints)
	fmt.Println("Your base statistics can't go beyond 15 ")
	input := printInputRequest("strength", &availablePoints, 5)
	str = input
	input = printInputRequest("dexterity", &availablePoints, 4)
	dex = input
	input = printInputRequest("constitution", &availablePoints, 3)
	con = input
	input = printInputRequest("intelligence", &availablePoints, 2)
	inte = input
	input = printInputRequest("wisdom", &availablePoints, 1)
	wis = input
	input = printInputRequest("charisma", &availablePoints, 0)
	cha = input

	goal = 50

	return
}

func printInputRequest(statName string, availablePoints *int, remainingStatToChose int) int {
	fmt.Printf("Please enter the number of points your want to add to your %s \n", statName)
	fmt.Printf("Remaining points : %v\n", *availablePoints)
	fmt.Printf("Remaining stats to choose : %v\n", remainingStatToChose)
	var input int
	fmt.Scan(&input)

	*availablePoints -= input
	return input
}
