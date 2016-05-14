package problems

import "github.com/johnmcdnl/projectEuler/utils"


//Summation of primes
/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
 */
func Problem010() int {
	return utils.SumSlice(utils.GetPrimesUntil(2000000))
}
