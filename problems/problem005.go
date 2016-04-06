package problems

import (
	"math"
)

func Problem005() int {
	targetNumDivisors := 500
	triangleNum := 0
	nextNaturalNumber := 1
	numFactors := 0

	for numFactors < targetNumDivisors {
		triangleNum = triangleNum + nextNaturalNumber
		nextNaturalNumber++
		numFactors = 2

		sqrt := int(math.Ceil(math.Sqrt(float64(triangleNum))))

		if (triangleNum / sqrt == sqrt) {
			numFactors = 1
		}

		for i := 2; i < sqrt; i++ {
			if (triangleNum % i == 0) {
				numFactors = numFactors + 2
			}

		}

	}
	return triangleNum
}

