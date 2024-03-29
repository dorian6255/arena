package main

import (
	"fmt"
	"reflect"
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
	player, err := getPlayerStats()
}

func getPlayerStats() (player, error) {
	baseCharacter := newPlayer()
	availablePoints := MAX_ALLO_STATS
	fmt.Println("Your character starts with 8 points in every of those stastistics :")
	fmt.Println("STRENGTH DEXTERITY CONSTITUTION INTELLIGENCE WISDOM AND CHARISMA")
	fmt.Println("You will now be asked to allocate %s points to your statistics", availablePoints)
	fmt.Println("Your base statistics can't go beyond 15 ")
	for i := 0; i < reflect.TypeOf(&baseCharacter).NumField(); i++ {
		field := reflect.TypeOf(&baseCharacter).Field(i)
		printInputRequest(field.Name, availablePoints)
	}
	return baseCharacter, nil
}

func printInputRequest(statName string, availablePoints int) (int, int, error) {
	fmt.Println("Please enter the number of points your want to add to your %s", statName)
	fmt.Println("Remaining points : %s", availablePoints)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil || input > 7 || input > availablePoints {
		return 0, 0, err
	}
	availablePoints -= 7
	return input, availablePoints, nil
}
