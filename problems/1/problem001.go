package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)

/**
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
 */
func Problem001() int {
	maxNum := 1000
	sum := 0
	for i := 1; i < maxNum; i++ {
		if (utils.IsDivisor(i, 3) || utils.IsDivisor(i, 5)) {
			sum += i
		}
	}
	return sum
}

