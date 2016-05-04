package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)


//Distinct primes factors

/*
The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
 */

func Problem047() int {

	inRowFound := 0
	target := 4

	primeFactors := [][]int{}
	for i := 0; i <= target; i++ {
		primeFactors = append(primeFactors, utils.RemoveDuplicateInts(utils.GetPrimeFactors(i)))
	}

	i := target
	for inRowFound < target {
		primeFactors = append(primeFactors, []int{})
		i++

		firstPrime := getFirstPrimeFactor(i);
		divideVal := i / firstPrime;

		if (divideVal == 1) {
			primeFactors[i] = []int{firstPrime}
			inRowFound = 0
		} else {
			updatedSlice := append(primeFactors[firstPrime], primeFactors[divideVal]...)
			primeFactors[i] = utils.RemoveDuplicateInts(updatedSlice)
			if len(primeFactors[i]) == target {
				inRowFound++
			} else {
				inRowFound = 0
			}
		}

	}
	return i - target + 1
}

func getFirstPrimeFactor(n int) int {
	if n % 2 == 0 {
		return 2
	}
	for i := 3; i <= n; i = i + 2 {
		for (n % i == 0) {
			return i
		}
	}

	return 0
}