package main

import (
	"fmt"
	"log"
	"os"
)

type player struct {
	STR  int
	DEX  int
	CON  int
	INTE int
	WIS  int
	CHA  int
}

func newPlayer() player {
	return player{
		STR:  8,
		DEX:  8,
		CON:  8,
		INTE: 8,
		WIS:  8,
		CHA:  8,
	}
}

const MAX_ALLO_STATS = 27 //Based on baldur's gate 3 character creator

func main() {
	fmt.Println("Welcome to the Arena")
	player := getPlayerStats()
	if err != nil {
		log.Fatal(err)
	}
	roundNumber := 1
	for {
		if ok := fightRound(player, roundNumber); !ok {
			break
		}
		roundNumber++
	}

}

func getPlayerStats() player {
	baseCharacter := newPlayer()
	availablePoints := MAX_ALLO_STATS
	fmt.Println("Your character starts with 8 points in every of those stastistics :")
	fmt.Println("STRENGTH DEXTERITY CONSTITUTION INTELLIGENCE WISDOM AND CHARISMA")
	fmt.Println("You will now be asked to allocate %s points to your statistics", availablePoints)
	fmt.Println("Your base statistics can't go beyond 15 ")
	input := printInputRequest("strength", &availablePoints)
	baseCharacter.STR += input
	input = printInputRequest("dexterity", &availablePoints)
	baseCharacter.DEX += input
	input = printInputRequest("constitution", &availablePoints)
	baseCharacter.CON += input
	input = printInputRequest("intelligence", &availablePoints)
	baseCharacter.INTE += input
	input = printInputRequest("wisdom", &availablePoints)
	baseCharacter.WIS += input
	input = printInputRequest("charisma", &availablePoints)
	baseCharacter.CHA += input
	return baseCharacter
}

func printInputRequest(statName string, availablePoints *int) int {
	fmt.Println("Please enter the number of points your want to add to your %s", statName)
	fmt.Println("Remaining points : %s", availablePoints)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil || input > 7 || input > *availablePoints {
		log.Fatal(err)
		os.Exit(2)
	}
	*availablePoints -= input
	return input
}

func fightRound(player player, roundNumber int) bool {

}
