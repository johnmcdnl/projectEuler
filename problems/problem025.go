package problems

import (
	"math/big"
)

func Problem025() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	index := big.NewInt(0)

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	for a.Cmp(&limit) < 0 {
		index.Add(index, big.NewInt(1))

		a.Add(a, b)
		a, b = b, a
	}
	return index
}
