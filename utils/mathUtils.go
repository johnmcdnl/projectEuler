package utils

import (
	"math"
)

func IsDivisor(number int, divisor int) bool {
	return number % divisor == 0
}

func IsPrime(number int) bool {
	if (number == 1) {
		return false
	} else if (number < 4) {
		return true
	} else if (number % 2 == 0) {
		return false
	} else if (number < 9) {
		return true
	} else if (number % 3 == 0) {
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

/**
 Retrieve all prime numbers as far as limit using Sieve of Eratosthenes
 */
func GetPrimesUntil(limit int) []int {

	composite := make([]bool, limit)
	composite[0] = true
	composite[1] = true
	composite[2] = false

	primes := []int{2}
	lastPrimeFound := 2

	j := 0
	for lastPrimeFound <= limit {

		for i := lastPrimeFound; i < limit; i += lastPrimeFound {
			composite[i] = true
		}
		composite[lastPrimeFound] = false

		for j = lastPrimeFound + 1; j < limit; j++ {
			if !composite[j] {
				lastPrimeFound = j
				break
			}
		}
		if j >= limit {
			return primes
		}
		primes = append(primes, lastPrimeFound)
	}

	return primes
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

func GetPrimeFactors(n int) []int {
	pfs := []int{}

	for i := 2; i <= n; i++ {
		for (n % i == 0) {
			pfs = append(pfs, i)
			n /= i;
		}
	}

	return pfs
}

func GenerateNFibonacciValues(num int) []int {
	a := 1
	b := 2
	fibs := []int{1, 2}
	for (len(fibs) < num) {
		a, b = b, a + b
		fibs = append(fibs, b)
	}
	return fibs
}

