package main

import (
	"flag"
	"fmt"

	"go-core.course/math/pkg/fibonacci"
)

func main() {
	inputNumber, err := parseInputNumber(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Высчитываем числo Фибоначчи для", inputNumber)

	fibonacciResult := fibonacci.CalculateNumber(inputNumber)
	fmt.Println("Число Фибоначчи: ", fibonacciResult)
}

func parseInputNumber(maxInputNumber int) (int, error) {
	inputNumber := flag.Int("n", -1, "Значение для вычисления числа Фибоначчи")
	flag.Parse()

	if *inputNumber > maxInputNumber {
		return *inputNumber, fmt.Errorf("Максимальное значение входного параметра - %d", maxInputNumber)
	}

	return *inputNumber, nil
}
