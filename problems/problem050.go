package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)


//Consecutive prime sum  -- https://projecteuler.net/problem=50

/*
The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
 */

var max50 = 1000000

func Problem050() int {
	primes := utils.GetPrimesUntilSum(max50)

	var primeSubArray [][]int
	longest := 0

	var longestArray []int

	for i := 0; i < len(primes); i++ {
		for j := i; j <= len(primes); j++ {
			if j - i > longest {
				subArray := primes[i:j]
				if isValidSubArrayBelowMax(subArray) {
					if len(subArray) > longest {
						longest = len(subArray)
						longestArray = subArray
					}
					primeSubArray = append(primeSubArray, subArray)
				}
			}
		}
	}

	return utils.SumSlice(longestArray)
}



func isValidSubArrayBelowMax(arr []int) bool {

	sum := 0

	for _, val := range arr {
		sum += val
		if sum > max50 {
			return false
		}
	}

	return utils.IsPrime(sum)
}
