package main

import (
	"strconv"
	"strings"

	"github.com/AfonsoTaborda/AdventOfCode2022/utils"
	"github.com/sirupsen/logrus"
)

func getFullyContainedRangesCount(input []string) (int, int) {
	fullyContainedRangesCount := 0
	overlappingRangesCount := 0

	for i := 0; i < len(input); i++ {
		first := strings.Split(input[i], ",")[0]
		second := strings.Split(input[i], ",")[1]

		firstLower, err := strconv.ParseInt(strings.Split(first, "-")[0], 10, 64)
		if err != nil {
			logrus.Fatal(err)
		}

		firstUpper, err := strconv.ParseInt(strings.Split(first, "-")[1], 10, 64)
		if err != nil {
			logrus.Fatal(err)
		}

		secondLower, err := strconv.ParseInt(strings.Split(second, "-")[0], 10, 64)
		if err != nil {
			logrus.Fatal(err)
		}

		secondUpper, err := strconv.ParseInt(strings.Split(second, "-")[1], 10, 64)
		if err != nil {
			logrus.Fatal(err)
		}

		// Count fully contained ranges for each elf pair's assignments
		if (firstLower <= secondLower && firstUpper >= secondUpper) || (secondLower <= firstLower && secondUpper >= firstUpper) {
			logrus.Printf("Found fully contained pair! First Elf: %v; Second Elf: %v", first, second)
			fullyContainedRangesCount++
		}

		// Count overlapping ranges for each elf pair's assignments
		if (firstLower <= secondUpper && firstUpper >= secondLower) || (secondLower <= firstUpper && secondUpper >= firstLower) {
			logrus.Printf("Found overlapping pair! First Elf: %v; Second Elf: %v", first, second)
			overlappingRangesCount++
		}
	}

	return fullyContainedRangesCount, overlappingRangesCount
}

func main() {
	logrus.Println("Starting Day4!")
	fullData := utils.GetInputData("./Input/input.txt")
	logrus.Printf("Full dataset is of size: %v\n", len(fullData))
	fullyContainedRangesCount, overlappingRangesCount := getFullyContainedRangesCount(fullData)
	logrus.Printf("The number of pair assignments that are fully contained withing each other is: %v, and for overlapping ranges is: %v", fullyContainedRangesCount, overlappingRangesCount)
}
