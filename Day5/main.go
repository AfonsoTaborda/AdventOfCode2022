package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AfonsoTaborda/AdventOfCode2022/utils"
	"github.com/sirupsen/logrus"
)

type Stack struct {
	items []string
}

func (stack *Stack) push(item string) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() (item string) {
	topItem := len(stack.items) - 1
	item = stack.items[topItem]
	stack.items = stack.items[:topItem]
	return item
}

func (stack *Stack) addToBottom(item string) {
	stack.items = append([]string{item}, stack.items...)
}

func parseColumns(delimiter string) []int {
	columnsRaw := strings.Split(delimiter, "   ")
	var columns []int

	for i := 0; i < len(columnsRaw); i++ {
		columnsRaw[i] = strings.Replace(columnsRaw[i], " ", "", -1)
		columnNumber, err := strconv.ParseInt(columnsRaw[i], 10, 64)
		if err != nil {
			logrus.Fatal(err)
		}

		columns = append(columns, int(columnNumber))
	}

	return columns
}

func generateStacks(input []string) (stacks []Stack) {
	delimiter := " 1   2   3   4   5   6   7   8   9 "
	columns := parseColumns(delimiter)

	logrus.Printf("Parsed columns: %v", columns)

	stacks = make([]Stack, len(columns))

	for i := 0; i < len(input)-1; i++ {
		line := input[i]

		if line != delimiter && !strings.Contains(line, "move") {
			for j, r := range line {
				if r != '[' && r != ']' && r != ' ' {
					stacks[j/4].addToBottom(string(r))
				}
			}
		}
	}

	return stacks
}

func makeMoves(input []string, stacks []Stack) []Stack {
	var toMove, from, to int

	for _, line := range input {
		fmt.Sscanf(line, "move %d from %d to %d", &toMove, &from, &to)

		logrus.Debugf("toMove: %d, from: %d, to: %d", toMove, from, to)

		for i := 0; i < toMove; i++ {
			stacks[to-1].push(stacks[from-1].pop())
		}
	}

	return stacks
}

func makeManyMoves(input []string, stacks []Stack) []Stack {
	var toMove, from, to int

	for _, line := range input {
		fmt.Sscanf(line, "move %d from %d to %d", &toMove, &from, &to)

		logrus.Debugf("toMove: %d, from: %d, to: %d", toMove, from, to)

		tempStack := make([]Stack, len(stacks))

		for i := 0; i < toMove; i++ {
			tempStack[to-1].addToBottom(stacks[from-1].pop())
		}

		for i, stack := range tempStack {
			for _, item := range stack.items {
				stacks[i].push(item)
			}
		}
	}

	return stacks
}

func findTopItem(stacks []Stack) []Stack {
	result := make([]Stack, len(stacks))

	for i, stack := range stacks {
		result[i].push(stack.items[len(stack.items)-1])
	}

	return result
}

func arrangeContainers(input []string) (stacks []Stack) {
	stacks = generateStacks(input)

	logrus.Printf("Parsed initial state from the input data: %v", stacks)

	movedStacks := makeMoves(input, stacks)

	logrus.Printf("The first arragement of the Elve's containers is: %v", movedStacks)

	result := makeManyMoves(input, generateStacks(input))

	return findTopItem(result)
}

func main() {
	logrus.Println("Starting Day 5!")
	fullData := utils.GetInputData("./Input/input.txt")
	logrus.Printf("Full dataset is of size: %v\n", len(fullData))
	logrus.Printf("The final arragement of the Elve's containers is: %v", arrangeContainers(fullData))
}
