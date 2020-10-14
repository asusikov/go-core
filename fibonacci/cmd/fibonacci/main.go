package main

import (
	"flag"
	"fmt"

	"go-core.course/fibonacci/pkg/fibonacci"
)

func main() {
	var inputNumber = flag.Int("n", -1, "Значение для вычисления числа Фибоначчи")
	flag.Parse()
	fmt.Println("Высчитываем число Фибоначчи для", *inputNumber)

	var fibonacciResult, err = fibonacci.Calculate(*inputNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Число Фибоначчи равно", fibonacciResult)
}
