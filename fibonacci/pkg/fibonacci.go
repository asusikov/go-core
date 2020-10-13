package fibonacci

import (
  "fmt"
)

const maxInputNumber = 20

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
  if number > maxInputNumber {
    return fmt.Errorf("Максимальное значение входного параметра - %d", maxInputNumber)
  }
  return nil
}
