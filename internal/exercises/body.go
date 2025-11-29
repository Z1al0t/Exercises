package exercises

import (
	"exercises/pkg/text"
	"exercises/pkg/utility"
	"fmt"
	"strconv"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

// var Yellow = "\033[33m"
var Blue = "\033[34m"

//var Magenta = "\033[35m"
//var Cyan = "\033[36m"
//var Gray = "\033[37m"
//var White = "\033[97m"

var (
	SolvedTasksList []int
	input           = 10
	startingTime    = time.Time{}
)

//var taskStartingTime = time.Time{}

func Run() {
	startingTime = time.Now()
	maxIndex := maxExerciseRow() + 1
	SolvedTasksList = make([]int, 0, maxIndex)
	SolvedTasksList = append(SolvedTasksList, 0)
	taskIndex := 0
	moreExircises := true

	for {
		fmt.Println()
		fmt.Println("-------------------------------------------")
		fmt.Printf(text.Print.GetExercise)

		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Printf("-------------------------------------------")

		fmt.Println()
		if input == 0 {
			printSolvedExercises(SolvedTasksList)
			fmt.Printf(Blue+text.Print.SessionTimeAmount+"\n"+Reset, time.Since(startingTime).Truncate(time.Second).String())
			fmt.Println(text.Print.Terminate)
			time.Sleep(15 * time.Second)

			break
		} else if input == 1 {
			//taskStartingTime = time.Now()
			SolvedTasksList, taskIndex, moreExircises = utility.RandomExerciseIndex(maxIndex, SolvedTasksList)

			if moreExircises == false {
				fmt.Println()
				fmt.Println("-------------------------------")
				fmt.Println("      " + Red + text.Print.Congratulation + Reset)
				fmt.Println(Red + text.Print.AllExercisesSolved + Reset)
				fmt.Println("-------------------------------")
				fmt.Println()
				fmt.Println(text.Print.Terminate)
				time.Sleep(15 * time.Second)
				break
			}

			indexString := strconv.Itoa(taskIndex)
			randomExerciseFromExcel(indexString)

		} else {
			fmt.Println(text.Print.InvalidInput)
			continue

		}

	}

}
