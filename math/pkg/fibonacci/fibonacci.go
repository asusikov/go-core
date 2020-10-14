package fibonacci

import (
	"fmt"
)

const maxInputNumber = 20

func Calculate(deepSize int) (result []int, err error) {
	err = validateInputNumber(deepSize)
	if err != nil {
		return []int{}, err
	}

	if deepSize == 1 {
		result = []int{1}
	} else {
		result = calculate(deepSize)
	}
	return result, nil
}

func calculate(deepSize int) []int {
	first := 1
	second := 1
	return append(
		[]int{first, second},
		recursiveCalculate(first, second, deepSize-2)...,
	)
}

func recursiveCalculate(first int, second int, deepSize int) []int {
	if deepSize <= 0 {
		return []int{}
	}

	newElement := first + second
	return append(
		[]int{newElement},
		recursiveCalculate(second, newElement, deepSize-1)...,
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
