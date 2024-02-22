/**
	Summary:
	Create a program that can perform dice rolls. The program should report the results as detailed in the requirements.

	Requirements:
	* Print the sum of the dice roll
	* Print additional information in these circumstances:
	- "Snake eyes": when the total roll is 2, and total dice is 2
	- "Lucky 7": when the total roll is 7
	- "Even": when the total roll is even
	- "Odd": when the total roll is odd
	* The program must handle any number of dice, rolls, and sides
**/

package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var play = true

	for play {
		fmt.Println("Nombre de dés souhaité :")

		var dices int
		_, err := fmt.Scanln(&dices)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Nombre de lancer souhaité :")

		var rolls int
		_, err = fmt.Scanln(&rolls)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Nombre de côtés souhaité pour un dé :")

		var sides int
		_, err = fmt.Scanln(&sides)
		if err != nil {
			log.Fatal(err)
		}

		total := 0

		for roll := 0; roll < rolls; roll++ {
			rollTotal := 0
			fmt.Printf("\nLancé de dé n°%d:\n", roll+1)

			for dice := 0; dice < dices; dice++ {
				diceResult := randInt(1, sides)
				fmt.Printf("Dé n°%d: %d\n", dice+1, diceResult)
				rollTotal += diceResult
			}

			if roll == 2 && rollTotal == 2 {
				fmt.Println("Snake eyes !")
			} else if rollTotal == 7 {
				fmt.Println("Lucky 7")
			}

			fmt.Printf("Résultat du lancé n°%d: %d\n\n", roll+1, rollTotal)
			total += rollTotal
		}

		fmt.Printf("Total: %d\n", total)

		if total == 7 {
			fmt.Printf("Lucky 7\n")
		}

		if total%2 == 0 {
			fmt.Printf("Even\n\n")
		} else {
			fmt.Printf("Odd\n\n")
		}

		var playAgain int
		fmt.Println("Souhaitez-vous recommencer? 1 oui, 2 non")
		_, err = fmt.Scanln(&playAgain)
		if err != nil {
			log.Fatal(err)
		}

		if playAgain != 1 {
			play = false
		}
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
