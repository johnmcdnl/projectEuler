package problems

import (
	"math/big"
	"strconv"
)

//Factorial digit sum

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
 */
func Problem020() int {
	n := int64(100)
	sum := big.NewInt(1)

	for i := int64(1); i < n; i++ {
		sum = sum.Mul(sum, big.NewInt(i))
	}
	asBytes := []byte(sum.String())
	sumChars := 0

	for _,bytes := range asBytes{
		asNum, _ := strconv.Atoi(string(bytes))
		sumChars +=asNum
	}
	return sumChars
}