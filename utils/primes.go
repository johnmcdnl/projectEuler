package utils

func GetPrimeFactors(n int) []int {
	pfs := []int{}

	for i := 2; i <= n; i++ {
		for (n % i == 0) {
			pfs = append(pfs, i)
			n /= i;
		}
	}

	return pfs
}

func GetPrimesUntilSum(sum int) []int {
	return getPrimesUntil(sum, sum)
}

/**
 Retrieve all prime numbers as far as limit using Sieve of Eratosthenes
 */
func GetPrimesUntil(limit int) []int {
	return getPrimesUntil(limit, 0)
}

func getPrimesUntil(limit int, maxSum int) []int {

	composite := make([]bool, limit)
	composite[0] = true
	composite[1] = true
	composite[2] = false

	primes := []int{2}
	lastPrimeFound := 2

	j := 0
	for lastPrimeFound <= limit {

		for i := lastPrimeFound; i < limit; i += lastPrimeFound {
			composite[i] = true
		}
		composite[lastPrimeFound] = false

		for j = lastPrimeFound + 1; j < limit; j++ {
			if !composite[j] {
				lastPrimeFound = j
				break
			}
		}
		if j >= limit {
			return primes
		}
		primes = append(primes, lastPrimeFound)
		if maxSum != 0 {
			if SumSlice(primes) >= maxSum {
				return primes
			}
		}
	}

	return primes
}

func IsPrime(number int) bool {
	if (number == 1) {
		return false
	} else if (number < 4) {
		return true
	} else if (number % 2 == 0) {
		return false
	} else if (number < 9) {
		return true
	} else if (number % 3 == 0) {
		return false
	}

	f := 5
	//All primes greater than 3 can be written in the form 6k+/-1.
	for f * f <= number {
		if (number % 5 == 0) {
			return false
		}
		if ( number % (f + int(2) ) == 0) {
			return false
		}
		f = f + int(6)
		if (number % f == 0) {
			return false
		}
	}

	return true
}