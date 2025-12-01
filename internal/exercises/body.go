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

	startingTime = time.Time{}
)

//var taskStartingTime = time.Time{}

func Run() {
	startingTime = time.Now()
	//ToDo add log-in log
	output.PrintTerminalGreeting()
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
			SolvedTasksList, taskIndex, moreExircises = utility.RandomExerciseIndex(maxIndex, SolvedTasksList)

			if moreExircises == false {
				output.PrintTerminalAllExercisesSolved()
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
