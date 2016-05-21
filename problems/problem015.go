package problems

//Lattice paths
/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?

 */
func Problem015() int {
	gridSize := 20
	pascalRow := PascalTriangleRow(gridSize*2)
	return pascalRow[gridSize]
}

func PascalTriangleRow(n int) []int {

	row := make( []int, n + 1, n + 1 )
	row[0], row[n] = 1, 1

	for i := 0; i < int(n / 2) ; i++ {
		x := row[ i ] * (n - i) / (i + 1)
		row[ i + 1 ], row[ n - 1 - i ] = x, x
	}

	return row
}
