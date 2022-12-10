package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetInputData(inputDataPath string) []string {
	inputFile, err := os.Open(inputDataPath)
	if err != nil {
		logrus.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var fullList []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fullList = append(fullList, line)
	}

	inputFile.Close()

	logrus.Println("Finished reading the input file...")

	return fullList
}

func FindMaxIntValue(input []int) int {
	max := input[0]

	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}

func UniqueStringBytes(input []string) []string {
	result := []string{}

	for i := 0; i < len(input); i++ {
		temp := []byte{}
		for j := 0; j < len(input[i]); j++ {
			if len(temp) != 0 {
				if !strings.Contains(string(temp), string(input[i][j])) {
					temp = append(temp, input[i][j])
				}
			} else {
				temp = append(temp, input[i][j])
			}
		}
		result = append(result, string(temp[:]))
	}

	logrus.Printf("The resulting unique string array is: %v\n", result)

	return result
}
