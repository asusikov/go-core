package main

import (
	"flag"
	"fmt"

	"go-core.course/math/pkg/fibonacci"
)

func main() {
	var inputNumber = parseInputNumber()

	var fibonacciResult, err = fibonacci.Calculate(*inputNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Число Фибоначчи равно", fibonacciResult)
}

func parseInputNumber() *int {
	var inputNumber = flag.Int("n", -1, "Значение для вычисления числа Фибоначчи")
	flag.Parse()
	fmt.Println("Высчитываем число Фибоначчи для", *inputNumber)
	return inputNumber
}
