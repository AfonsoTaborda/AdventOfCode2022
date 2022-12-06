package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getInputData() []int {
	fmt.Println("Started getting data...")

	inputFile, err := os.Open("./Input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	var linesInt []int

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	inputFile.Close()

	for _, line := range lines {

		if line == "" {
			linesInt = append(linesInt, -1)
		} else {
			parsedInt, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			linesInt = append(linesInt, int(parsedInt))
		}
	}

	fmt.Println("Finished reading the input file...")

	return linesInt
}

func findMaxIntValue(input []int) int {
	max := input[0]

	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}

func findTopThreeElvesTotal(input []int) []int {
	// Start by finding the sum of what the top 3 elves are carrying
	var list []int

	sort.Ints(input)

	for i := 0; i < 3; i++ {
		list = append(list, input[len(input)-i-1])
		fmt.Println("Top " + strconv.Itoa(i) + " elf carries: " + strconv.Itoa(input[len(input)-i-1]))
	}

	return list
}

func calculateTotalCarriedByElf(rawInput []int) []int {
	elfCount := 0
	sum := 0
	var allElfLoads []int

	for _, amount := range rawInput {
		if amount != -1 {
			sum += amount
		} else {
			allElfLoads = append(allElfLoads, sum)
			fmt.Println("Elf " + strconv.Itoa(elfCount) + ", carries: " + strconv.Itoa(sum))
			sum = 0
			elfCount += 1
		}
	}

	return allElfLoads
}

func main() {
	fmt.Println("Go Day 1 Started !!!")

	lines := getInputData()

	totalCarriedByElf := calculateTotalCarriedByElf(lines)

	elfCarryingTheMost := findMaxIntValue(totalCarriedByElf)

	fmt.Printf("The maximum carried by an elf is: %v\n", elfCarryingTheMost)

	topThreeCarryingElves := findTopThreeElvesTotal(totalCarriedByElf)

	sumTopThree := topThreeCarryingElves[0] + topThreeCarryingElves[1] + topThreeCarryingElves[2]

	fmt.Printf("The top 3 carrying elves are carrying in total: %v\n", sumTopThree)
}
