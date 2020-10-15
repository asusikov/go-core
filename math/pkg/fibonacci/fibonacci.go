package fibonacci

const firstFibonacciElement = 1
const secondFibonacciElement = 1

func Calculate(deepSize int) []int {
	var result []int
	switch {
	case deepSize < 1:
		result = []int{}
	case deepSize == 1:
		result = []int{firstFibonacciElement}
	default:
		result = calculate(deepSize)
	}
	return result
}

func calculate(deepSize int) []int {
	return append(
		[]int{firstFibonacciElement, secondFibonacciElement},
		recursiveCalculate(firstFibonacciElement, secondFibonacciElement, deepSize-2)...,
	)
}

func recursiveCalculate(first int, second int, deepLevel int) []int {
	if deepLevel <= 0 {
		return []int{}
	}

	newElement := first + second
	return append(
		[]int{newElement},
		recursiveCalculate(second, newElement, deepLevel-1)...,
	)
}
