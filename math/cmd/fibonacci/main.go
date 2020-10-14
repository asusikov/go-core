package main

import (
	"flag"
	"fmt"

	"go-core.course/math/pkg/fibonacci"
)

func main() {
	inputNumber := parseInputNumber()

	fibonacciResult, err := fibonacci.Calculate(*inputNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Числа Фибоначчи: ", fibonacciResult)
}

func parseInputNumber() *int {
	inputNumber := flag.Int("n", -1, "Значение для вычисления чисел Фибоначчи")
	flag.Parse()
	fmt.Println("Высчитываем числа Фибоначчи для", *inputNumber)
	return inputNumber
}
