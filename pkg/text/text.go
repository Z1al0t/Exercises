package text

import (
	"encoding/json"
	"exercises/internal/config"
	"fmt"
	"os"
	"path/filepath"
)

type DisplayText struct {
	Greeting             string `json:"greeting"`
	GetExercise          string `json:"get_exercise"`
	ExerciseTime         string `json:"exercise_time"`
	InvalidInput         string `json:"invalid_input"`
	Congratulation       string `json:"congratulation"`
	SolvedExerciseAmount string `json:"solved_exercise_amount"`
	EnterName            string `json:"enter_name"`
	SessionTimeAmount    string `json:"session_time_amount"`
	Terminate            string `json:"terminate"`
	AllExercisesSolved   string `json:"all_exercises_solved"`
	ExerciseNumber       string `json:"exercise_number"`
}

var Print DisplayText

func GetTextFromFile(text *DisplayText) {
	file, err := os.ReadFile(filepath.Join(config.PrintingFilePath, config.PrintingFileName))
	if err != nil {
		fmt.Printf("Error reading from %s file.\n", config.PrintingFileName)
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(file, &text)
	if err != nil {
		fmt.Printf("Error unmarshalling from %s file.\n", config.PrintingFileName)
		fmt.Println(err)
		return
	}

}
