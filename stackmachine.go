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
		num, _ := strconv.Atoi(command)

		// if err != nil {
		// 	return errors.New("invalid command encountered: not a number and not included in list of allowed operations")
		// }

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

func (m *Machine) isStackEmpty() bool {
	if len(m.stack) == 0 {
		return true
	}
	return false
}

func (m *Machine) notEnoughNumsInStack(minimum int) bool {
	if len(m.stack) < minimum {
		return true
	}
	return false
}

func (m *Machine) processDUPOperation() {
	lastNumInStack := m.getNumFromBackOfStack(1)
	m.appendToStack(lastNumInStack)
}

func (m *Machine) processPOPOperation() {
	m.removeElementsFromStack(1)
}

func (m *Machine) processAddOperation() error {
	lastNumInStack, secondLastNumInStack := m.getNumsForDoubleOps()
	numTotal := lastNumInStack + secondLastNumInStack

	if isNumberOutOfBounds(numTotal) {
		return errors.New("stack integer add overflow")
	}

	m.cleanupAfterDoubleOps(numTotal)
	return nil
}

func (m *Machine) processMinusOperation() error {
	lastNumInStack, secondLastNumInStack := m.getNumsForDoubleOps()
	result := lastNumInStack - secondLastNumInStack

	if isNumberOutOfBounds(result) {
		return errors.New("stack integer underflow")
	}

	m.cleanupAfterDoubleOps(result)
	return nil
}

func (m *Machine) processMultiplyOperation() error {
	lastNumInStack, secondLastNumInStack := m.getNumsForDoubleOps()
	result := lastNumInStack * secondLastNumInStack

	if isNumberOutOfBounds(result) {
		return errors.New("stack integer overflow")
	}

	m.cleanupAfterDoubleOps(result)
	return nil
}

func (m *Machine) processSUMOperation() error {
	var sumTotal int
	for _, num := range m.stack {
		sumTotal += num
		if isNumberOutOfBounds(sumTotal) {
			return errors.New("stack integer sum overflow")
		}
	}

	m.removeElementsFromStack(len(m.stack))
	m.appendToStack(sumTotal)
	return nil
}

func (m *Machine) appendToStack(num int) {
	m.stack = append(m.stack, num)
}

func (m *Machine) getNumFromBackOfStack(inverseIndex int) int {
	return m.stack[len(m.stack)-inverseIndex]
}

func (m *Machine) getNumsForDoubleOps() (int, int) {
	return m.stack[len(m.stack)-1], m.stack[len(m.stack)-2]
}

func (m *Machine) removeElementsFromStack(numberOfElements int) {
	m.stack = m.stack[:len(m.stack)-numberOfElements]
}

func (m *Machine) cleanupAfterDoubleOps(result int) {
	m.removeElementsFromStack(2)
	m.appendToStack(result)
}
