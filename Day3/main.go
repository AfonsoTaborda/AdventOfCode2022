package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AfonsoTaborda/AdventOfCode2022/utils"
)

func getInputData() ([]string, []string, []string) {
	inputFile, err := os.Open("./Input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var firstHalfList []string
	var secondHalfList []string
	var fullList []string

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()

		firstHalfList = append(firstHalfList, rucksack[:len(rucksack)/2])

		secondHalfList = append(secondHalfList, rucksack[len(rucksack)/2:])

		fullList = append(fullList, rucksack)
	}

	inputFile.Close()

	fmt.Println("Finished reading the input file...")

	return firstHalfList, secondHalfList, fullList
}

func getPriorities() map[byte]int {
	lowerCase := byte('a')
	upperCase := byte('A')
	var priorities = make(map[byte]int)

	// Calculate lowercase priorities
	for i := 1; i < 27; i++ {
		priorities[lowerCase] = i
		lowerCase++
	}

	// Calculate uppercase priorities
	for i := 27; i < 53; i++ {
		priorities[upperCase] = i
		upperCase++
	}

	fmt.Printf("Generated lowercase and uppercase priorities: %v\n", priorities)

	return priorities
}

func getPrioritiesByBadgeGroup(inputList []string) int {
	priorities := getPriorities()
	prioritiesScoreSum := 0

	for i := 0; i < len(inputList); i += 3 {
		var first = inputList[i]
		var second = inputList[i+1]
		var third = inputList[i+2]

		for j := 0; j < len(first); j++ {
			if strings.Contains(second, string(first[j])) {
				if strings.Contains(third, string(first[j])) {
					prioritiesScoreSum += priorities[first[j]]
					continue
				}
			}
		}
	}

	return prioritiesScoreSum
}

func getPrioritiesSum(firstHalfList []string, secondHalfList []string) int {
	priorities := getPriorities()
	prioritiesScoreSum := 0

	for i := 0; i < len(firstHalfList); i++ {
		for j := 0; j < len(firstHalfList[i]); j++ {
			if strings.Contains(secondHalfList[i], string(firstHalfList[i][j])) {
				prioritiesScoreSum += priorities[firstHalfList[i][j]]
			}
		}
	}

	return prioritiesScoreSum
}

func main() {
	log.Println("Starting Day 3!")
	firstHalfList, secondHalfList, fullList := getInputData()

	fmt.Printf("Got first Half %v and the second half %v of the rucksack\n", string(firstHalfList[0]), string(secondHalfList[0]))

	prioritiesScoreSum := getPrioritiesSum(utils.UniqueStringBytes(firstHalfList), utils.UniqueStringBytes(secondHalfList))

	log.Printf("The sum of the similar rucksack items is: %v", prioritiesScoreSum)

	badgePrioritySum := getPrioritiesByBadgeGroup(utils.UniqueStringBytes(fullList))

	log.Printf("The sum of the similar rucksack items by badge groups is: %v", badgePrioritySum)
}
