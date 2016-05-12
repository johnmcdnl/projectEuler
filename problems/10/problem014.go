package problems

//Longest Collatz sequence

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
 */

var m = make(map[int]int)

func Problem014() int {

	target := 1000000
	for i := 1; i < target; i++ {
		m[i] = getCollatzSequenceLength(i)
	}
	longest := 0
	longestStart := 0
	for index, length := range m {
		if length > longest {
			longest = length
			longestStart = index
		}
	}
	return longestStart
}

func getCollatzSequenceLength(n int) int {
	var seq []int
	for n != 1 {
		n = getNextCollatzValue(n)
		seq = append(seq, n)
		if lengthPreCalc, ok := m[n]; ok {
			return len(seq) + lengthPreCalc
		}

	}
	return len(seq)
}

func getNextCollatzValue(n int) int {
	if n % 2 == 0 {
		n /= 2
	} else {
		n = 3 * n + 1
	}
	return n
}