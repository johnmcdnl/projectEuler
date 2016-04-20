package problems



/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
 */
func Problem004() int {

	var multiples []int

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			multiples = append(multiples, i * j)
		}
	}

	var largestPalindrome int
	for _, multiple := range multiples {
		if isPalindrome(string(multiple)) {
			if multiple > largestPalindrome {
				largestPalindrome = multiple
			}
		}
	}

	return largestPalindrome
}

func isPalindrome(s string) bool {
	r := Reverse(s)
	return r == s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}