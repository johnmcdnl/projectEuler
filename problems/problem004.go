package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
	"strconv"
)

/**
Problem 4 - Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
 */
func Problem004() int {
	largestPalindrome := 0

	for i := 999; i >= 0; i-- {
		if (largestPalindrome / i > i) {
			break
		}
		for j := 999; j >= 0; j-- {
			if (largestPalindrome / j > i) {
				break
			}
			if utils.IsPalindrome(strconv.Itoa(i * j)) && i * j > largestPalindrome {
				largestPalindrome = i * j
			}
		}
	}

	return largestPalindrome
}

