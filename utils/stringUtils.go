package utils

func IsPalindrome(s string) bool {
	bytes := []byte(s)

	a := 0
	z := len(bytes) - 1
	for a <= len(bytes) / 2 {
		if bytes[a] != bytes[z] {
			return false
		}
		a++
		z--
	}

	return true
}