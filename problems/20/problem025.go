package problems

import (
	"math/big"
)

func Problem025() int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	index := 0

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	for a.Cmp(&limit) < 0 {
		index++

		a.Add(a, b)
		a, b = b, a
	}
	return index
}