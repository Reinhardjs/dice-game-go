package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

func removeElementAt(array *[]int, index int) {
	arr := *array
	copy(arr[index:], arr[index+1:]) // Shift array[i+1:] left one index.
	arr[len(arr)-1] = -1             // Erase last element (write zero value).
	*array = arr[:len(*array)-1]     // Truncate slice.
}

func evaluate(playersDices *[][]int, playersScore []int) {
	newPlayerDices := make([][]int, len(*playersDices))

	for playerIndex, dices := range *playersDices {
		var newDices []int

		for i := 0; i < len((*playersDices)[playerIndex]); i++ {
			if dices[i] == 6 || dices[i] == 1 {
				playersScore[playerIndex] += 1
			} else {
				newDices = append(newDices, dices[i])
			}
		}

		newPlayerDices[playerIndex] = newDices
	}

	for playerIndex, dices := range *playersDices {

		for i := 0; i < len((*playersDices)[playerIndex]); i++ {
			if dices[i] == 1 {
				var nextPlayer int

				if playerIndex+1 >= len(*playersDices) {
					nextPlayer = 0
				} else {
					nextPlayer = playerIndex + 1
				}

				newPlayerDices[nextPlayer] = append(newPlayerDices[nextPlayer], 1)
			}
		}
	}

	for playerIndex, _ := range *playersDices {
		(*playersDices)[playerIndex] = newPlayerDices[playerIndex]
	}
}

func printCurrentState(playersDices [][]int, playersScore []int) {
	for index, dices := range playersDices {
		dicesJson, _ := json.Marshal(dices)
		fmt.Println("Player #" + strconv.Itoa(index+1) + " (" + strconv.Itoa(playersScore[index]) + "): " + string(dicesJson))
	}
}

func isGameEnd(playersDices [][]int) bool {
	playersWithDiceCount := len(playersDices)
	for _, dices := range playersDices {
		if len(dices) == 0 {
			playersWithDiceCount -= 1
		}
	}

	return playersWithDiceCount == 1
}

func getRandomFrom(min int, toMax int) int {
	return rand.Intn(toMax) + min
}

func main() {
	fmt.Println("This game is automatic Dice game")
	fmt.Println("You just need to enter the number of players and number of dices")

	fmt.Println("Enter number of players: ")
	var numOfPlayers int
	fmt.Scanln(&numOfPlayers)

	fmt.Println("Enter number of dices : ")
	var numOfDices int
	fmt.Scanln(&numOfDices)

	var playersDices [][]int = make([][]int, numOfPlayers)
	var playersScore []int = make([]int, numOfPlayers)

	fmt.Println("First round")
	for i := range playersDices {
		playersDices[i] = make([]int, numOfDices)

		for j := range playersDices[i] {
			playersDices[i][j] = getRandomFrom(1, 6)
		}
	}
	printCurrentState(playersDices, playersScore)
	evaluate(&playersDices, playersScore)
	fmt.Println("After evaluate : ")
	printCurrentState(playersDices, playersScore)

	fmt.Println()
	fmt.Println()

	round := 2
	for !isGameEnd(playersDices) {
		fmt.Println("Round " + strconv.Itoa(round))
		for i := range playersDices {
			playersDices[i] = make([]int, numOfDices)

			for j := 0; j < len(playersDices[i]); j++ {
				playersDices[i][j] = getRandomFrom(1, 6)
			}
		}
		printCurrentState(playersDices, playersScore)
		evaluate(&playersDices, playersScore)
		fmt.Println("After evaluate : ")
		printCurrentState(playersDices, playersScore)
		fmt.Println()
		fmt.Println()

		round++
	}
}
