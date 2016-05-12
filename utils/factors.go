package utils

import "math"

func LowestCommonMultiple(x, y int) int {
	return (x * y) / GreatestCommonDivisor(x, y)
}

func GreatestCommonDivisor(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func GetAllFactors(num int) []int {

	factors := []int{1, num}
	sqrt := int(math.Ceil(math.Sqrt(float64(num))))

	for i := 2; i < sqrt; i++ {
		if (num % i == 0) {
			factors = append(factors, i)
			factors = append(factors, num / i)
		}
	}
	if (num % sqrt == 0) {
		factors = append(factors, sqrt)
	}

	return factors
}

func GetNumberOfDivisors(number int) int {
	return len(GetAllFactors(number))
}