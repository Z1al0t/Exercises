package output

import (
	"exercises/pkg/text"
	"fmt"
	"strings"
	"time"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"

	//  Yellow = "\033[33m"
	Blue = "\033[34m"

	// Magenta = "\033[35m"
	// Cyan = "\033[36m"
	// Gray = "\033[37m"
	White = "\033[97m"
)

func getInput(lower bool) string {
	var input string
	fmt.Printf(White)
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		panic(err)
	}
	fmt.Printf(Reset)
	if input == "" {
		fmt.Println("Input cannot be blank")
		panic(err)
	}

	input = strings.Replace(input, "\n", "", -1)
	input = strings.TrimSpace(input)
	if lower {
		input = strings.ToLower(input)
	}
	return input
}

func PrintTerminalGreeting() {
	// Greeting message structure
	fmt.Println("-------------------------------------------")
	fmt.Println()
	fmt.Println(Green + text.Print.Greeting + Reset)
	fmt.Println()
	fmt.Println("-------------------------------------------")

}

func PrintTerminalGetExercise() string {
	// Get Exercises menu
	fmt.Println()
	fmt.Println("-------------------------------------------")
	fmt.Printf(text.Print.GetExercise)
	userInput := getInput(true)
	fmt.Printf("-------------------------------------------")
	fmt.Println()
	return userInput
}

func PrintTerminalExit(solved []int, st time.Time) {
	// Exit message
	amount := len(solved) - 1
	fmt.Printf(Blue+text.Print.SolvedExerciseAmount+"\n"+Reset, amount)
	fmt.Printf(Blue+text.Print.SessionTimeAmount+"\n"+Reset, time.Since(st).Truncate(time.Second).String())
	fmt.Println(text.Print.Terminate)
	time.Sleep(15 * time.Second)
}

func PrintTerminalAllExercisesSolved() {
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println("      " + Red + text.Print.Congratulation + Reset)
	fmt.Println(Red + text.Print.AllExercisesSolved + Reset)
	fmt.Println("-------------------------------")
	fmt.Println()
	fmt.Println(text.Print.Terminate)
	time.Sleep(15 * time.Second)
}

func PrintTerminalExerciseNumber(number string) {
	fmt.Println()
	fmt.Println(Green + "-------------" + Reset)
	fmt.Printf(Green+text.Print.ExerciseNumber+"\n"+Reset, number)
	fmt.Println(Green + "-------------" + Reset)
}

func PrintTerminalGetName(name *string) {
	fmt.Println(text.Print.EnterName)
	*name = getInput(false)
}

func PrintTerminalGetSolvingStatus(solved *string, st time.Time, exercise string) {
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Printf(text.Print.ExerciseSolvingStatus)
	*solved = getInput(true)
	fmt.Println("-------------------------------")

	fmt.Printf(Blue+text.Print.SolvedExerciseNumber+Reset, exercise, *solved)
	timeSpend := time.Since(st).Truncate(time.Second).String()
	fmt.Printf(Blue+text.Print.SolvedExerciseTime+Reset, timeSpend)
}
