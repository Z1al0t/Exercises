package exercises

import (
	"exercises/internal/output"
	//"exercises/pkg/logs"
	"exercises/pkg/text"
	"exercises/pkg/utility"
	"fmt"
	"strconv"
	"time"
)

var (
	SolvedTasksList []int

	startingTime     = time.Time{}
	taskStartingTime = time.Time{}
	name             string
)

func Run() {
	startingTime = time.Now()
	//ToDo add log-in log
	output.PrintTerminalGreeting()

	output.PrintTerminalGetName(&name)

	maxIndex := maxExerciseRow() + 1
	SolvedTasksList = make([]int, 0, maxIndex)
	SolvedTasksList = append(SolvedTasksList, 0)
	taskIndex := 0
	moreExircises := true

	for {
		var userInput string

		userInput = output.PrintTerminalGetExercise()

		if userInput == "exit" {
			//Todo log log-out
			output.PrintTerminalExit(SolvedTasksList, startingTime)
			break
		} else if userInput == "get" {
			//Todo log starting exercise time
			taskStartingTime = time.Now()
			SolvedTasksList, taskIndex, moreExircises = utility.RandomExerciseIndex(maxIndex, SolvedTasksList)

			if moreExircises == false {
				output.PrintTerminalAllExercisesSolved()
				break
			}

			indexString := strconv.Itoa(taskIndex)
			randomExerciseFromExcel(indexString)

			exerciseSolvedStatus := ""
			output.PrintTerminalGetSolvingStatus(&exerciseSolvedStatus, taskStartingTime, indexString)

		} else if userInput == "help" {
			//Todo help menu
			fmt.Println("This feature is not active at the moment, sorry.")

		} else {
			fmt.Println(text.Print.InvalidInput)
			continue

		}

	}

}
