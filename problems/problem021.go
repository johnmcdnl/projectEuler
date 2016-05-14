package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)

//Amicable numbers
/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
 */
func Problem021() int {

	sumDivisors := make(map[int]int)

	for i := 1; i < 10000; i++ {
		sumDivisors[i] = utils.SumSlice(utils.RemoveDuplicateInts(utils.GetAllFactors(i))) - i
	}
	amicablePairs := make(map[int]int)

	for i := 0; i < len(sumDivisors); i++ {
		for j := 0; j < len(sumDivisors); j++ {
			//If d(i) = j and d(j) = i, where i ≠ j
			if i != j {
				if sumDivisors[i] == j {
					if sumDivisors[j] == i {
						amicablePairs[i] = j
					}
				}
			}
		}
	}
	sumPairs := 0
	for key := range amicablePairs {
		sumPairs += key
	}

	return sumPairs
}
