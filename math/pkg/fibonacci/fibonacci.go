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
	result = append([]int{1, 1}, recursiveCalculate(1, 1, deepSize - 2)...)
	return result, nil
}

func recursiveCalculate(first int, second int, deepSize int) []int {
	if deepSize == 0 {
		return []int{}
	}
	newElement := first + second
	return append([]int{newElement}, recursiveCalculate(second, newElement, deepSize - 1)...)
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
