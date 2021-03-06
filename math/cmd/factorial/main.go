package main

import (
	"flag"
	"fmt"

	"go-core.course/math/pkg/factorial"
)

func main() {
	inputNumber := parseInputNumber()

	factorialResult, err := factorial.Calculate(inputNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Факториал равен", factorialResult)
}

func parseInputNumber() int {
	inputNumber := flag.Int("n", -1, "Значение для вычисления факториала")
	flag.Parse()
	fmt.Println("Высчитываем факториал для", *inputNumber)
	return *inputNumber
}
