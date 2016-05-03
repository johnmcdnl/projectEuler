package problems

import "github.com/johnmcdnl/projectEuler/utils"


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

	target := 4
	inRowFound := 0
	i := 1
	for inRowFound < target {
		unique := utils.RemoveDuplicateInts(utils.GetPrimeFactors(i))
		if len(unique) == target {
			inRowFound++
		} else {
			inRowFound = 0
		}
		i++
	}

	return i - target
}



