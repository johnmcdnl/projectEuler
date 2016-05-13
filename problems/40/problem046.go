package problems

//Goldbach's other conjecture

/**
It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×1^2
15 = 7 + 2×2^2
21 = 3 + 2×3^2
25 = 7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
 */
import (
	"github.com/johnmcdnl/projectEuler/utils"
	"math"
)

var start = 33
var primes []int = utils.GetPrimesUntil(start)

func Problem046() int {
	i := start
	for {
		if utils.IsPrime(i) {
			primes = append(primes, i)
			i += 2
		} else {
			if isGoldbachOtherConjecture(i) {
				i += 2
			} else {
				return i
			}
		}
	}
}

func isGoldbachOtherConjecture(n int) bool {

	for _, prime := range primes {
		subtract := n - prime
		if subtract % 2 == 0 {
			subtract /= 2
			maxSqrt := int(math.Ceil(math.Sqrt(float64(subtract))))
			for i := maxSqrt; i >= 1; i-- {
				if i * i == subtract {
					return true
				}
			}
		}
	}
	return false
}