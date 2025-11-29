package exercises

import (
	"fmt"
	"strconv"
	"time"
)

var input = 10

func Run() {
	maxIndex := maxRow() + 1
	solvedTasks := make([]int, 0, maxIndex)
	solvedTasks = append(solvedTasks, 0)
	taskIndex := 0
	moreExircises := true

	for {
		fmt.Println()
		fmt.Println("-------------------------------------------")
		fmt.Printf("Type 0 to exit, 1 to get random exercise:")

		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Printf("-------------------------------------------")

		fmt.Println()
		if input == 0 {
			printSolvedExercises(solvedTasks)
			fmt.Println("This program will terminate in 15 seconds.")
			time.Sleep(15 * time.Second)

			break
		} else if input == 1 {
			solvedTasks, taskIndex, moreExircises = randomIndex(maxIndex, solvedTasks)
			if moreExircises == false {
				fmt.Println()
				fmt.Println("-------------------------------")
				fmt.Println("      Congratulations!")
				fmt.Println("All exercises have been solved.")
				fmt.Println("-------------------------------")
				fmt.Println()
				fmt.Println("This program will terminate in 15 seconds.")
				time.Sleep(15 * time.Second)
				break
			}

			indexString := strconv.Itoa(taskIndex)
			randomExercise(indexString)

		} else {
			fmt.Println("Invalid input")
			continue

		}

	}

}
