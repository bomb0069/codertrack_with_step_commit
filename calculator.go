package codetrack

import (
	"strconv"

	"github.com/golang-collections/collections/stack"
)

func add(firstOperand string, secondOperand string) string {
	firstStack := stack.New()
	secondStack := stack.New()
	resultStack := stack.New()

	for _, charactor := range firstOperand {
		firstStack.Push(int(charactor) - 48)
	}

	for _, charactor := range secondOperand {
		secondStack.Push(int(charactor) - 48)
	}

	todd := 0
	firstInt := 0
	secondInt := 0
	result := 0

	for firstStack.Len() != 0 || secondStack.Len() != 0 {

		if firstStack.Len() != 0 {
			firstInt = firstStack.Pop().(int)
		} else {
			firstInt = 0
		}

		if secondStack.Len() != 0 {
			secondInt = secondStack.Pop().(int)
		} else {
			secondInt = 0
		}

		result = todd + firstInt + secondInt

		if result >= 10 {
			todd = 1
			result = result - 10
		} else {
			todd = 0
		}

		resultStack.Push(result)
	}

	if todd == 1 {
		resultStack.Push(todd)
	}

	resultString := ""
	for resultStack.Len() != 0 {
		resultString = resultString + strconv.Itoa(resultStack.Pop().(int))
	}

	return resultString
}
