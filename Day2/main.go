package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	ElfRock       byte = 'A'
	ElfPaper      byte = 'B'
	ElfScissors   byte = 'C'
	Rock          byte = 'X'
	Paper         byte = 'Y'
	Scissors      byte = 'Z'
	Loss          int  = 0
	Draw          int  = 3
	Win           int  = 6
	RockScore     int  = 1
	PaperScore    int  = 2
	ScissorsScore int  = 3
)

func getRawData() ([]byte, []byte) {
	inputFile, err := os.Open("./Input/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var yourInput []byte
	var elfInput []byte

	for fileScanner.Scan() {
		tempElfInput := fileScanner.Text()[0]
		tempYourInput := fileScanner.Text()[2]
		yourInput = append(yourInput, tempYourInput)
		elfInput = append(elfInput, tempElfInput)
	}

	inputFile.Close()

	fmt.Println("Finished reading the input file...")

	return yourInput, elfInput
}

func resolveToPoints(yourInput []byte, elfInput []byte) int {
	var score int = 0
	fmt.Printf("Computing score for %v total matches...\n", len(yourInput))
	for i := 0; i < len(yourInput); i++ {
		elfCurrent := elfInput[i]
		current := yourInput[i]
		switch elfCurrent {
		case ElfRock:
			if current == Scissors {
				score += Loss + ScissorsScore
			}
			if current == Paper {
				score += Win + PaperScore
			}
			if current == Rock {
				score += Draw + RockScore
			}
		case ElfPaper:
			if current == Scissors {
				score += Win + ScissorsScore
			}
			if current == Paper {
				score += Draw + PaperScore
			}
			if current == Rock {
				score += Loss + RockScore
			}
		case ElfScissors:
			if current == Scissors {
				score += Draw + ScissorsScore
			}
			if current == Paper {
				score += Loss + PaperScore
			}
			if current == Rock {
				score += Win + RockScore
			}
		default:
			log.Fatal("Error: Bad input data!")
			break
		}
	}
	return score
}

func resolveToPointsPartTwo(yourInput []byte, elfInput []byte) int {
	var score int = 0
	fmt.Printf("Computing score for %v total matches...\n", len(yourInput))
	for i := 0; i < len(yourInput); i++ {
		elfCurrent := elfInput[i]
		current := yourInput[i]
		switch current {
		case Rock:
			if elfCurrent == ElfScissors {
				score += Loss + PaperScore
			}
			if elfCurrent == ElfPaper {
				score += Loss + RockScore
			}
			if elfCurrent == ElfRock {
				score += Loss + ScissorsScore
			}
		case Paper:
			if elfCurrent == ElfScissors {
				score += Draw + ScissorsScore
			}
			if elfCurrent == ElfPaper {
				score += Draw + PaperScore
			}
			if elfCurrent == ElfRock {
				score += Draw + RockScore
			}
		case Scissors:
			if elfCurrent == ElfScissors {
				score += Win + RockScore
			}
			if elfCurrent == ElfPaper {
				score += Win + ScissorsScore
			}
			if elfCurrent == ElfRock {
				score += Win + PaperScore
			}
		default:
			log.Fatal("Error: Bad input data!")
			break
		}
	}
	return score
}

func calculateTotalScores(yourInput []byte, elfInput []byte) {
	totalScore := resolveToPoints(yourInput, elfInput)
	log.Printf("The total possible score with the given input data hints is: %v\n", totalScore)
	totalGuidedScore := resolveToPointsPartTwo(yourInput, elfInput)
	log.Printf("The total score following the Part 2 elf's guide is: %v\n", totalGuidedScore)
}

func main() {
	log.Println("Go Day 2 Started !!!")
	yourInput, elfInput := getRawData()
	fmt.Printf("All data has been gathered, the first line from your input is: %v\n", string(yourInput[0]))
	fmt.Printf("All data has been gathered, the first line from the Elf's input is: %v\n", string(elfInput[0]))
	calculateTotalScores(yourInput, elfInput)
}
