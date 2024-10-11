package main

import (
	"errors"
	"strconv"
	"strings"
)

const MAX_NUMBER_LIMIT = 50000

type Machine struct {
	stack []int
}

func StackMachine(commands string) (int, error) {
	if len(commands) == 0 {
		return 0, errors.New("empty command")
	}

	commandSlice := strings.Split(commands, " ")
	machine := Machine{
		stack: []int{},
	}

	for _, command := range commandSlice {
		err := processCommand(command, &machine)

		if err != nil {
			return 0, err
		}
	}

	firstNumInStack := machine.stack[0]
	return firstNumInStack, nil
}

func isNumberOutOfBounds(num int) bool {
	return num < 0 || num > MAX_NUMBER_LIMIT
}

func processCommand(command string, m *Machine) error {
	switch command {
	case "+":
		if len(m.stack) < 2 {
			return errors.New("not enough numbers in stack for add")
		}

		m.processAddOperation()

	default:
		num, err := strconv.Atoi(command)

		if err != nil {
			return errors.New("invalid command encountered: not a number and not included in list of allowed operations")
		}

		if isNumberOutOfBounds(num) {
			return errors.New("number out of bounds")
		}

		m.appendToStack(num)
	}

	return nil
}

func (m *Machine) processAddOperation() error {
	lastNumInStack := m.popFromStack(1)
	secondLastNumInStack := m.popFromStack(2)

	numTotal := lastNumInStack + secondLastNumInStack

	if isNumberOutOfBounds(numTotal) {
		return errors.New("stack integer overflow")
	}

	m.removeElementsFromStack(2)
	m.appendToStack(numTotal)
	return nil
}

func (m *Machine) appendToStack(num int) {
	m.stack = append(m.stack, num)
}

func (m *Machine) popFromStack(inverseIndex int) int {
	return m.stack[len(m.stack)-inverseIndex]
}

func (m *Machine) removeElementsFromStack(numberOfElements int) {
	m.stack = m.stack[:len(m.stack)-numberOfElements]
}
