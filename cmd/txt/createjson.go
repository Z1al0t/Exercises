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
		Greeting:             "Welcome to the Exercises app!",
		GetExercise:          "Type 0 to exit, 1 to get random exercise: ",
		ExerciseTime:         "You spent %v for this exercise.",
		InvalidInput:         "Invalid input!",
		Congratulation:       "Congratulation!",
		SolvedExerciseAmount: "You solved %d exercises at this session",
		SessionTimeAmount:    "You spent %v practicing exercises today",
		Terminate:            "This program will terminate in 15 seconds.",
		AllExercisesSolved:   "All exercises have been solved.",
		ExerciseNumber:       "Exercise #%s.",
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
