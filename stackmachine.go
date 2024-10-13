package main

import (
	"errors"
	"strconv"
	"strings"
)

const MAX_NUMBER_LIMIT = 50000
const MIN_ELEMENTS_FOR_SINGLE_OPS = 1
const MIN_ELEMENTS_FOR_DOUBLE_OPS = 2

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

	if machine.isStackEmpty() {
		return 0, nil
	}

	firstNumInStack := machine.stack[0]
	return firstNumInStack, nil
}

func processCommand(command string, m *Machine) error {
	switch command {
	case "POP":
		if m.notEnoughNumsInStack(MIN_ELEMENTS_FOR_SINGLE_OPS) {
			return errors.New("not enough nums in stack for a POP")
		}
		m.processPOPOperation()

	case "DUP":
		if m.notEnoughNumsInStack(MIN_ELEMENTS_FOR_SINGLE_OPS) {
			return errors.New("not enough nums in stack for a DUP")
		}
		m.processDUPOperation()

	case "CLEAR":
		stackLength := len(m.stack)
		m.removeElementsFromStack(stackLength)

	case "SUM":
		if m.isStackEmpty() {
			return errors.New("cannot SUM on empty stack")
		}
		return m.processSUMOperation()

	case "+":
		if m.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
			return errors.New("not enough numbers in stack for add")
		}
		return m.processAddOperation()

	case "-":
		if m.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
			return errors.New("not enough numbers in stack for minus")
		}
		return m.processMinusOperation()

	case "*":
		if m.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
			return errors.New("not enough numbers in stack for multiply")
		}
		return m.processMultiplyOperation()

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

func isNumberOutOfBounds(num int) bool {
	return num < 0 || num > MAX_NUMBER_LIMIT
}
