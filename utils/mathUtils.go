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

	var factors int
	sqrt := int(math.Ceil(math.Sqrt(float64(number))))

	for i := 2; i < sqrt; i++ {
		if (number % i == 0) {
			factors = factors + 2
		}
	}
	return factors
}

