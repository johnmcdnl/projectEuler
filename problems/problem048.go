package problems

import (
	"math/big"
	"strconv"
)

//Self Powers
/*
The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.s
 */
func Problem048() int {

	sum := big.NewInt(int64(0))
	for i:=int64(1); i<=1000; i++{
		pow := new(big.Int).Exp(big.NewInt(i), big.NewInt(i), nil)
		sum.Add(sum, pow)
	}
	asBytes := []byte(sum.String())
	last10 := asBytes[len(asBytes)-10:]
	ans, _ := strconv.Atoi(string(last10))
	return ans
}



