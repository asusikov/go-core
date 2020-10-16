package fibonacci

func CalculateNumber(number int) int {
	switch number {
	case 0, 1:
		return number
	default:
		return CalculateNumber(number - 2) + CalculateNumber(number - 1)
	}
}
