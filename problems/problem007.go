package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
	"math"
)

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10,001st prime number?
 */
func Problem007() int {
	nthPrime := 10001

	piTimes := float64(nthPrime) * math.Log(float64(nthPrime))
	primes := utils.GetPrimesUntil(int(piTimes))

	primeCounter := len(primes)
	i := primes[primeCounter - 1]

	if primeCounter >= nthPrime {
		return primes[nthPrime - 1]
	}

	for primeCounter < nthPrime {
		i++
		if (utils.IsPrime(i)) {
			primeCounter++
		}
	}

	return i
}