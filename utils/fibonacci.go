package utils

func GenerateFibonacciUntil(max int) []int {

	a := 1
	b := 2
	fibs := []int{1, 2}
	for (b < max) {
		a, b = b, a + b
		fibs = append(fibs, b)
	}
	return fibs[:len(fibs) - 1]

}

func GenerateNFibonacciValues(num int) []int {
	a := 1
	b := 2
	fibs := []int{1, 2}
	for (len(fibs) < num) {
		a, b = b, a + b
		fibs = append(fibs, b)
	}
	return fibs
}