package main

import (
	"github.com/AfonsoTaborda/AdventOfCode2022/utils"
	"github.com/sirupsen/logrus"
)

func getMarkerIndices(input []string, differenCharacters int) map[rune]int {
	markers := make(map[rune]int)

	for _, lines := range input {
		for index, char := range lines {
			tempArr := make(map[byte]bool)

			for k := 0; k < differenCharacters; k++ {
				tempArr[lines[index+k]] = true
			}

			if len(tempArr) == differenCharacters {
				markers[char] = index + differenCharacters
				break
			}
		}
	}

	return markers
}

func main() {
	logrus.Println("Starting Day 6!")
	fullData := utils.GetInputData("./Input/input.txt")
	logrus.Printf("Founr different characters: %v", getMarkerIndices(fullData, 4))
	logrus.Printf("Fourteen different markers: %v", getMarkerIndices(fullData, 14))
}
