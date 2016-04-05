package problems

import "github.com/johnmcdnl/projectEuler/utils"

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10,001st prime number?
 */
func Problem007() uint64 {
	nthPrime := 10001
	primeCounter := 0
	i := uint64(0)
	for primeCounter < nthPrime {
		i++
		if (mathUtils.IsPrime(i)) {
			primeCounter++
		}
	}

	return i
}