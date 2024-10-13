package main

import (
	"errors"
)

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
	lastNumInStack, secondLastNumInStack := m.getNumbersForDoubleOps()
	numTotal := lastNumInStack + secondLastNumInStack

	if isNumberOutOfBounds(numTotal) {
		return errors.New("stack integer add overflow")
	}

	m.processOperation(numTotal, 2)
	return nil
}

func (m *Machine) processMinusOperation() error {
	lastNumInStack, secondLastNumInStack := m.getNumbersForDoubleOps()
	result := lastNumInStack - secondLastNumInStack

	if isNumberOutOfBounds(result) {
		return errors.New("stack integer underflow")
	}

	m.processOperation(result, 2)
	return nil
}

func (m *Machine) processMultiplyOperation() error {
	lastNumInStack, secondLastNumInStack := m.getNumbersForDoubleOps()
	result := lastNumInStack * secondLastNumInStack

	if isNumberOutOfBounds(result) {
		return errors.New("stack integer overflow")
	}

	m.processOperation(result, 2)
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

	stackLength := len(m.stack)
	m.processOperation(sumTotal, stackLength)
	return nil
}

func (m *Machine) appendToStack(num int) {
	m.stack = append(m.stack, num)
}

func (m *Machine) getNumFromBackOfStack(inverseIndex int) int {
	return m.stack[len(m.stack)-inverseIndex]
}

func (m *Machine) getNumbersForDoubleOps() (int, int) {
	return m.stack[len(m.stack)-1], m.stack[len(m.stack)-2]
}

func (m *Machine) removeElementsFromStack(numberOfElements int) {
	m.stack = m.stack[:len(m.stack)-numberOfElements]
}

func (m *Machine) processOperation(resultToAppend int, numOfNumsToRemove int) {
	m.removeElementsFromStack(numOfNumsToRemove)
	m.appendToStack(resultToAppend)
}
