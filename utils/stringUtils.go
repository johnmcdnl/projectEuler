package utils

func IsPalindrome(s string) bool {
	bytes := []byte(s)

	a := 0
	z := len(bytes)-1
	midPoint := len(bytes) / 2
	for a <= midPoint {
		if bytes[a] != bytes[z] {
			return false
		}
		a++
		z--
	}

	return true
}