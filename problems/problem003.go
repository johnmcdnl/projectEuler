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

	factors := utils.GetPrimeFactors(600851475143)
	largestPrimeFactor := 0

	for _, factor := range factors {
		largestPrimeFactor = utils.Max(largestPrimeFactor, factor)
	}

	return largestPrimeFactor
}
