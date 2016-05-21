package problems

import (
	"math/big"
	"strconv"
)
/*
Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
What is the sum of the digits of the number 2^1000?
 */
func Problem016() int {
	base := big.NewInt(2)
	power := big.NewInt(1000)
	asString := base.Exp(base, power, nil).String()

	sum := 0
	for _, char := range asString {
		asNum, _ := strconv.Atoi(string(char))
		sum += asNum
	}

	return sum
}
