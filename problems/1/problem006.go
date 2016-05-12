package problems

func Problem006() int {

	target := 100

	sumSquares := 0
	for i := 1; i <= target; i++ {
		sumSquares += (i * i)
	}

	squareSums := 0
	for i := 1; i <= target; i++ {
		squareSums += i
	}

	squareSums *= squareSums

	return squareSums - sumSquares
}