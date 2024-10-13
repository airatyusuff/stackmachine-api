package main

import (
	"testing"
)

func TestStartsWithEmptyStack(t *testing.T) {
	_, err := StackMachine("")

	if err == nil {
		t.Error("expected error due to no results")
	}
}

// Write your own TDD tests here as you develop
func TestCommandNotValid(t *testing.T) {
	commands := []string{
		"12+",
		"INVALID",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for invalid commands")
		}
	}
}
func TestHasNumberOutOfBounds(t *testing.T) {
	commands := []string{
		"50001",
		"-34",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for integer out of bounds")
		}
	}
}

func TestInvalidCommandsForAdd(t *testing.T) {
	commands := []string{
		"+",
		"12 +",
		"1 2 + +",
		"1 2 + 4 + +",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for add operation")
		}
	}
}

func TestInvalidCommandsForDUP(t *testing.T) {
	commands := []string{
		"DUP",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for DUP operation")
		}
	}
}

func TestInvalidCommandsForPOP(t *testing.T) {
	commands := []string{
		"POP",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for POP operation")
		}
	}
}

func TestInvalidCommandsForMinus(t *testing.T) {
	commands := []string{
		"12 -",
		"-",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for - operation")
		}
	}
}

func TestInvalidCommandsForMultiply(t *testing.T) {
	commands := []string{
		"12 *",
		"*",
	}

	for _, command := range commands {
		_, err := StackMachine(command)

		if err == nil {
			t.Error("expected error for * operation")
		}
	}
}

func TestClearCommandError(t *testing.T) {
	commands := []string{
		"12 CLEAR",
		"12 23 CLEAR",
	}

	for _, command := range commands {
		got, _ := StackMachine(command)

		if got != 0 {
			t.Errorf("expected %d, got %d", 0, got)
		}
	}
}

func TestSumCommandError(t *testing.T) {
	commands := []string{
		"SUM",
		"50000 45 SUM",
	}

	for _, command := range commands {
		got, err := StackMachine(command)

		if err == nil {
			t.Errorf("expected error for SUM operation got %d", got)
		}
	}
}

func TestCommandReturnsExpectedOutput(t *testing.T) {
	commands := []string{
		"501",
	}

	for _, command := range commands {
		got, _ := StackMachine(command)

		expected := 501
		if got != expected {
			t.Errorf("expected %d, but got %d", expected, got)
		}
	}
}

/*
All these tests must pass for completion
*/
func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		expectedErr error
	}{
		{name: "empty error", commands: "", expected: 0},
		{name: "add overflow", commands: "50000 DUP +", expected: 0},
		{name: "too few add", commands: "99 +", expected: 0},
		{name: "too few minus", commands: "99 -", expected: 0},
		{name: "too few multiply", commands: "99 *", expected: 0},
		{name: "empty stack", commands: "99 CLEAR", expected: 0},
		{name: "sum single value", commands: "99 SUM", expected: 99},
		{name: "sum empty", commands: "SUM", expected: 0},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5},
		{name: "single integer", commands: "9876", expected: 9876},
		{name: "invalid command", commands: "DOGBANANA", expected: 0},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60},
		{name: "minus", commands: "2 5 -", expected: 3},
		{name: "underflow minus", commands: "5 2 -", expected: 0},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000},
		{name: "overflow plus", commands: "50000 1 +", expected: 0},
		{name: "overflow single value", commands: "50001", expected: 0},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0},
		{name: "clear empties stack", commands: "CLEAR", expected: 0},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120},
	}

	for _, test := range tests {
		got, err := StackMachine(test.commands)

		if got > 0 && err != nil {
			t.Errorf("expected no error, got %s", err.Error())
		}
		if got != test.expected {
			t.Errorf("(%s) got %v, want %v", test.commands, got, test.expected)
		}
	}
}
