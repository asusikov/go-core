package main

import (
	"flag"
	"fmt"

	"go-core.course/fibonacci/pkg/fibonacci"
)

func main() {
	var inputNumber = flag.Int("n", -1, "help message for flag n")
	flag.Parse()
	fmt.Println("Высчитываем число Фибоначчи для", *inputNumber)

	var fibonacciResult, err = fibonacci.Calculate(*inputNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Число Фибонначи равно", fibonacciResult)
}
