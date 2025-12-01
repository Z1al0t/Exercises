package exercises

import (
	"exercises/internal/config"
	"exercises/pkg/text"
	"fmt"

	"github.com/xuri/excelize/v2"
)

// maxExerciseRow defines amount of Exercises in Excel file
func maxExerciseRow() int {

	file, err := excelize.OpenFile(config.ExercisesFileName)
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

// randomExercise prints random Exercise task
func randomExerciseFromExcel(index string) {

	file, err := excelize.OpenFile(config.ExercisesFileName)
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
					fmt.Println(Green + "-------------" + Reset)
					fmt.Printf(Green+text.Print.ExerciseNumber+"\n"+Reset, row[0])
					fmt.Println(Green + "-------------" + Reset)
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
