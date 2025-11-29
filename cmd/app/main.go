package main

import (
	"exercises/internal/exercises"
	"exercises/pkg/text"
)

func main() {

	text.GetTextFromFile(&text.Print)
	exercises.Run()

}
