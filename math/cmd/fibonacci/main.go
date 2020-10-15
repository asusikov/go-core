package main

import (
	"flag"
	"fmt"

	"go-core.course/math/pkg/fibonacci"
)

const maxInputNumber = 20

func main() {
	inputNumber, err := parseInputNumber()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Высчитываем числа Фибоначчи для", inputNumber)

	fibonacciResult := fibonacci.Calculate(inputNumber)
	fmt.Println("Числа Фибоначчи: ", fibonacciResult)
}

func parseInputNumber() (int, error) {
	inputNumber := flag.Int("n", -1, "Значение для вычисления чисел Фибоначчи")
	flag.Parse()

	if *inputNumber > maxInputNumber {
		return *inputNumber, fmt.Errorf("Максимальное значение входного параметра - %d", maxInputNumber)
	}

	return *inputNumber, nil
}
