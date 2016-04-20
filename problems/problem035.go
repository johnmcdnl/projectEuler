package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
	"fmt"
	"sort"
	"strconv"
)

/**
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
How many circular primes are there below one million?
https://projecteuler.net/problem=35
 */

var primes = []int{}

func Problem035() int {

	max := 100

	for i := 1; i < max; i++ {
		if mathUtils.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	for _, prime := range primes {
		// element is the element from someSlice for where we are
		if isCircularPrime(prime) {
			//fmt.Println(prime)
		}
	}

	fmt.Println(primes)
	return 0
}

func getCircularValues(number int) []int {

	chars := []byte(strconv.Itoa(number))
	fmt.Println(chars)
	return []int{1}
}

func isCircularPrime(number int) bool {
	circularValues := getCircularValues(number)
	for _, circular := range circularValues {
		if (sort.SearchInts(primes, circular) == len(circularValues)) {
			return false
		}
	}
	return true
}



/**
public static void permutation(String str) {
permutation("", str);
}

private static void permutation(String prefix, String str) {
int n = str.length();
if (n == 0) System.out.println(prefix);
else {
for (int i = 0; i < n; i++)
permutation(prefix + str.charAt(i), str.substring(0, i) + str.substring(i+1, n));
}
}
 */

