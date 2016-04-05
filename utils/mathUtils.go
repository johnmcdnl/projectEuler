package mathUtils

import (
	"math"
)

func IsDivisor(number int, divisor int) bool {
	return number % divisor == 0
}

/**
All primes greater than 3 can be written in the form 6k+/-1.
Any number n can have only one primefactor greater than n .
 */
func IsPrime(number uint64) bool {
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

	rootN := uint64(math.Floor(math.Sqrt(float64(number))))
	f := uint64(5)
	for f <= rootN {
		if (number % 5 == 0) {
			return false
		}
		if ( number % (f + uint64(2) ) == 0) {
			return false
		}
		f = f + uint64(6)
		if (number % f == 0) {
			return false
		}
	}

	return true
}

