package factorial

import (
	"fmt"
)

func Calculate(number int) (result int, err error) {
	err = validateInputNumber(number)
	if err != nil {
		return 0, err
	}
	return calculate(number), nil
}

func calculate(number int) int {
	if number == 1 {
		return 1
	}
	nextResult := calculate(number - 1)
	return number * nextResult
}

func validateInputNumber(number int) error {
	if number < 1 {
		return fmt.Errorf("Входной параметр должен быть больше или равен 1")
	}
	return nil
}
