package fibonacci

import (
	"fmt"
)

const maxInputNumber = 20
const firstFibonacciElement = 1
const secondFibonacciElement = 1

func Calculate(deepSize int) (result []int, err error) {
	err = validateInputNumber(deepSize)
	if err != nil {
		return []int{}, err
	}

	if deepSize == 1 {
		result = []int{firstFibonacciElement}
	} else {
		result = calculate(deepSize)
	}
	return result, nil
}

func calculate(deepSize int) []int {
	return append(
		[]int{firstFibonacciElement, secondFibonacciElement},
		recursiveCalculate(firstFibonacciElement, secondFibonacciElement, deepSize-2)...,
	)
}

func recursiveCalculate(first int, second int, deepLevel int) []int {
	if deepLevel <= 0 {
		return []int{}
	}

	newElement := first + second
	return append(
		[]int{newElement},
		recursiveCalculate(second, newElement, deepLevel-1)...,
	)
}

func validateInputNumber(number int) error {
	if number < 1 {
		return fmt.Errorf("Входной параметр должен быть больше или равен 1")
	}

	if number > maxInputNumber {
		return fmt.Errorf("Максимальное значение входного параметра - %d", maxInputNumber)
	}

	return nil
}
