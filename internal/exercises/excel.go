package exercises

import (
	"exercises/internal/config"
	"fmt"
	"math/rand"
	"slices"

	"github.com/xuri/excelize/v2"
)

// maxRow defines amount of Exercises in Excel file
func maxRow() int {

	file, err := excelize.OpenFile(config.FileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	result := 0

	for i, _ := range rows {
		result = i
	}

	return result
}

// randomIndex gets random exercise which wasn't solved yet
func randomIndex(m int, done []int) ([]int, int, bool) {
	moreExercises := true
	index := 0
	for i := 0; i < m; i++ {
		randIndex := rand.Intn(m)
		if slices.Contains(done, randIndex) {
			if len(done) == m {
				moreExercises = false
				break
			}
			i--
			continue
		}
		done = append(done, randIndex)
		index = randIndex
		break
	}
	return done, index, moreExercises
}

// randomExercise prints random Exercise task
func randomExercise(index string) {

	file, err := excelize.OpenFile(config.FileName)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	excelColumns := map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
		26: "AA",
		27: "AB",
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range rows {

		for j, cell := range row {
			if row[0] == index {
				if j == 0 {
					fmt.Println()
					fmt.Println("-------------")
					fmt.Printf("Exercise #%s.\n", row[0])
					fmt.Println("-------------")
					continue
				} else {
					column := excelColumns[j]

					header, err := file.GetCellValue("Sheet1", column+"1")
					if err != nil {
						fmt.Println(err)
					}
					fmt.Printf("%s:\n", header)
					fmt.Println(cell)
					continue
				}
			}

		}
	}

}

func printSolvedExercises(solved []int) {
	amount := len(solved) - 1
	fmt.Println("Solved Exercises:", amount)

}
