package mathUtils

import (
	"math"
)

func IsDivisor(number int, divisor int) bool {
	return number % divisor == 0
}

func IsPrime(number int) bool {
	if (number == 1) {
		return false
	}else if (number < 4) {
		return true
	}else if (number % 2 == 0) {
		return false
	}else if (number < 9) {
		return true
	}else if (number % 3 == 0) {
		return false
	}


	rootN := int(math.Floor(math.Sqrt(float64(number))))
	f := 5
	//All primes greater than 3 can be written in the form 6k+/-1.
	for f <= rootN {
		if (number % 5 == 0) {
			return false
		}
		if ( number % (f + int(2) ) == 0) {
			return false
		}
		f = f + int(6)
		if (number % f == 0) {
			return false
		}
	}

	return true
}

func GetNumberOfDivisors(number int) int {

	return len(GetAllFactors(number))
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


func GenerateFibonacciUntil(max int) []int {

	a := 1
	b := 2
	fibs := []int{1, 2}
	for (b < max) {
		a, b = b, a + b
		fibs = append(fibs, b)
	}
	return fibs[:len(fibs) - 1]

}

func GenerateNFibonacciValues(num int) []int{
	a := 1
	b := 2
	fibs := []int{1, 2}
	for (len(fibs) < num) {
		a, b = b, a + b
		fibs = append(fibs, b)
	}
	return fibs
}

