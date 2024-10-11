package main

import (
	"errors"
	"strconv"
	"strings"
)

const MAX_NUMBER_IN_STACK = 50000

type Store struct {
	stack    []int
	commands []string
}

func (s *Store) appendToStack(num int) {
	s.stack = append(s.stack, num)
}

func (s *Store) popFromStack(inverseIndex int) int {
	return s.stack[len(s.stack)-inverseIndex]
}

func (s *Store) removeElementsFromStack(numberOfElements int) {
	s.stack = s.stack[:len(s.stack)-numberOfElements]
}

func isNumberOutOfBounds(number int) bool {
	return number < 0 || number > MAX_NUMBER_IN_STACK
}

func StackMachine(commands string) (int, error) {
	if len(commands) == 0 {
		return 0, errors.New("")
	}

	machineStore := Store{
		stack:    []int{},
		commands: strings.Split(commands, " "),
	}

	for _, command := range machineStore.commands {
		switch command {
		case "DUP":
			if len(machineStore.stack) < 1 {
				return 0, errors.New("not enough numbers in stack for duplicating")
			}
			dupLastNumInStack := machineStore.popFromStack(1)
			machineStore.appendToStack(dupLastNumInStack)

		case "+":
			if len(machineStore.stack) < 2 {
				return 0, errors.New("not enough numbers in stack for add")
			}

			lastNumInStack := machineStore.popFromStack(1)
			secondLastNumInStack := machineStore.popFromStack(2)

			numTotal := lastNumInStack + secondLastNumInStack

			if isNumberOutOfBounds(numTotal) {
				return 0, errors.New("stack integer overflow")
			}

			machineStore.removeElementsFromStack(2)
			machineStore.appendToStack(numTotal)

		default:
			numberToAppendToStack, err := strconv.Atoi(command)
			if err != nil {
				return 0, errors.New("invalid command")
			}

			if isNumberOutOfBounds(numberToAppendToStack) {
				return 0, errors.New("number not within accepted range")
			}

			machineStore.appendToStack(numberToAppendToStack)
		}
	}

	result := machineStore.stack[0]
	return result, nil

}

// func isTotalWithinLimit(number int, secondnumber int) bool {
// 	total := number + secondnumber
// 	if total < 0 || total <= MAX_NUMBER_IN_STACK {
// 		return true
// 	}
// 	return false
// }

func main() {
	// main is unused - run using
	// go test ./...
}
