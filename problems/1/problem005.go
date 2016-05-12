package problems

/**
Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func Problem005() int {

	target := 20
	allPrimeFactors := false

	num := 1

	for !allPrimeFactors {
		allPrimeFactors = true
		for i := 1; i < target; i++ {
			if num % i != 0 {
				allPrimeFactors = false
				num++
				break
			}
		}
	}

	return num
}