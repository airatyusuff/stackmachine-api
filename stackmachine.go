package main

import (
	"errors"
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
		operation := CreateOperation(command, &machine)
		err := operation.Execute()

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

func CreateOperation(command string, m *Machine) Operation {
	switch command {
	case "POP":
		return &PopOperation{machine: m}

	case "DUP":
		return &DupOperation{machine: m}

	case "CLEAR":
		return &ClearOperation{machine: m}

	case "SUM":
		return &SumOperation{machine: m}

	case "+":
		return &AddOperation{machine: m}

	case "-":
		return &MinusOperation{machine: m}

	case "*":
		return &MultiplyOperation{machine: m}

	default:
		return &NumbersOperation{machine: m, args: command}
	}
}
