package main

import (
	"errors"
	"strconv"
)

type Operation interface {
	Execute() error
}
type PopOperation struct {
	machine *Machine
}
type DupOperation struct {
	machine *Machine
}
type ClearOperation struct {
	machine *Machine
}
type SumOperation struct {
	machine *Machine
}
type AddOperation struct {
	machine *Machine
}
type MinusOperation struct {
	machine *Machine
}
type MultiplyOperation struct {
	machine *Machine
}
type NumbersOperation struct {
	machine *Machine
	args    string
}

func (op *PopOperation) Execute() error {
	if op.machine.notEnoughNumsInStack(MIN_ELEMENTS_FOR_SINGLE_OPS) {
		return errors.New("not enough nums in stack for a POP")
	}
	op.machine.processPOPOperation()
	return nil
}

func (op *DupOperation) Execute() error {
	if op.machine.notEnoughNumsInStack(MIN_ELEMENTS_FOR_SINGLE_OPS) {
		return errors.New("not enough nums in stack for a DUP")
	}
	op.machine.processDUPOperation()
	return nil
}

func (op *ClearOperation) Execute() error {
	stackLength := len(op.machine.stack)
	op.machine.removeElementsFromStack(stackLength)
	return nil
}

func (op *SumOperation) Execute() error {
	if op.machine.isStackEmpty() {
		return errors.New("cannot SUM on empty stack")
	}
	return op.machine.processSUMOperation()
}

func (op *AddOperation) Execute() error {
	if op.machine.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
		return errors.New("not enough numbers in stack for add")
	}
	return op.machine.processAddOperation()
}

func (op *MinusOperation) Execute() error {
	if op.machine.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
		return errors.New("not enough numbers in stack for minus")
	}
	return op.machine.processMinusOperation()
}

func (op *MultiplyOperation) Execute() error {
	if op.machine.notEnoughNumsInStack(MIN_ELEMENTS_FOR_DOUBLE_OPS) {
		return errors.New("not enough numbers in stack for multiply")
	}
	return op.machine.processMultiplyOperation()
}

func (op *NumbersOperation) Execute() error {
	num, err := strconv.Atoi(op.args)
	if err != nil {
		return errors.New("invalid command encountered")
	}
	if isNumberOutOfBounds(num) {
		return errors.New("number out of bounds")
	}
	op.machine.appendToStack(num)
	return nil
}
