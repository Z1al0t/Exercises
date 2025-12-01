package main

import (
	"encoding/json"
	"exercises/pkg/text"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	jsonFileName := "text_eng.json" //replace with needed filename
	filePath := "pkg/text/"

	myText := text.DisplayText{
		Greeting:              "Welcome to the Exercises app!",
		GetExercise:           "     Type your command: \n<exit> to quit \n<help> to see all commands \n<get>  to get random exercise\nCommand: ",
		ExerciseTime:          "You spent %v for this exercise.",
		InvalidInput:          "Invalid input!",
		Congratulation:        "Congratulation!",
		SolvedExerciseAmount:  "You solved %d exercises at this session",
		SessionTimeAmount:     "You spent %v practicing exercises today",
		Terminate:             "This program will terminate in 15 seconds.",
		AllExercisesSolved:    "All exercises have been solved.",
		ExerciseNumber:        "Exercise #%s.",
		EnterName:             "Please enter your name:",
		ExerciseSolvingStatus: "Please enter exercise solving status in percentage:\n<100> for fully solved\n<0> for not solved\n<50> for solved only half\nSolved: ",
		SolvedExerciseNumber:  "Exercise #%s solved %s%%\n",
		SolvedExerciseTime:    "You spent %v for this exercise.",
	}

	jsonText, err := json.Marshal(myText)
	if err != nil {
		fmt.Println("Error marshaling to JSON.", err)
		return
	}
	err = os.WriteFile(filepath.Join(filePath, jsonFileName), jsonText, 0644)
	if err != nil {
		fmt.Println("Error creating file.", err)
		return
	}
	fmt.Println("Created file", jsonFileName)

}
