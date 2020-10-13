package fibonacci

import (
	"flag"
	"fmt"
)

func main() {
	var inputNumber = flag.Int("n", 0, "help message for flag n")
	flag.Parse()
	fmt.Println("Высчитываем число Фибоначчи для", *inputNumber)

	var fibonacciResult, err = fibonacci.Calculate(*inputNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Число Фибонначи равно", fibonacciResult)
}
