package main

import (
	"fmt"
)

 type player struct{
	STR int
	DEX int
	CON int
	INTE int
	WIS int
	CHA int
}

func newPlayer() Player {
	return Player{
		STR: 8 
		DEX: 8
		CON: 8
		INT: 8
		WIS: 8
		CHA: 8
	}
}

const MAX_ALLO_STATS = 27 //Based on baldur's gate 3 character creator

func main() {
	fmt.Println("Welcome to the Arena")

}

func getPlayerStats(firstTime bool) player,error {
	baseCharacter := newPlayer()
	availablePoints := MAX_ALLO_STATS 
	fmt.Println("Your character starts with 8 points in every of those stastistics :")
	fmt.Println("STRENGTH DEXTERITY CONSTITUTION INTELLIGENCE WISDOM AND CHARISMA")
	fmt.Println("You will now be asked to allocate %s points to your statistics",availablePoints)
	fmt.Println("Your base statistics can't go beyond 15 ")
	printInputRequest("strength")
	var str, dex, con, inte, wis, cha int
	_, err := fmt.Scan()
}

func printInputRequest(statName string) {
	fmt.Println("Please enter the number of points your want to add to your %s",statName)
}