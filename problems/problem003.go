package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)

/**
Problem 3 - Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
func Problem003() int {

	factors := mathUtils.GetAllFactors(600851475143)

	largestPrimeFactor := 0

	for _, factor := range factors {
		if mathUtils.IsPrime(factor) && factor > largestPrimeFactor {
			largestPrimeFactor = factor
		}
	}

	return largestPrimeFactor
}

