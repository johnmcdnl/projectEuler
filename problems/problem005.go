package problems

import "github.com/johnmcdnl/projectEuler/utils"

/**
Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func Problem005() int {

	target := 20
	result := 1

	for i := 1; i <= target; i++ {
		result = utils.LowestCommonMultiple(result, i)
	}

	return result
}