package problems

import (
	"strconv"
	"github.com/johnmcdnl/projectEuler/utils"
)

//Truncatable primes
/**
The number 3797 has an interesting property.
Being prime itself, it is possible to continuously remove digits from left to right,
and remain prime at each stage: 3797, 797, 97, and 7.
Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
 */

func Problem037() int {

	currentList := []string{"2", "3", "5", "7", "9"}
	appendValues := []string{"1", "3", "7", "9"}

	counter := 0
	truncatablePrimes := []int{}

	loop:
	for {
		currentList = getNextSetOfCandidates(currentList, appendValues)
		for _, value := range currentList {
			asNum, _ := strconv.Atoi(value);
			if utils.IsPrime(asNum) {
				if isTruncatablePrime(value) {
					counter++
					truncatablePrimes = append(truncatablePrimes, asNum)
				}
			}
			if counter == 11 {
				break loop
			}
		}
	}

	return utils.SumSlice(primes)
}

func isTruncatablePrime(number string) bool {

	for i := 1; i < len(number); i++ {
		if asNum, _ := strconv.Atoi(number[:len(number) - i]); !utils.IsPrime(asNum) {
			return false
		}
	}

	for i := 1; i < len(number); i++ {
		if asNum, _ := strconv.Atoi(number[i:]); !utils.IsPrime(asNum) {
			return false
		}
	}

	return true

}

func getNextSetOfCandidates(currentList, appendValues  []string) []string {
	newList := []string{}

	for _, current := range currentList {
		for _, appendVal := range appendValues {
			if current[len(current) - 1:] != appendVal {
				newList = append(newList, current + appendVal)
			}
		}
	}

	return newList
}
