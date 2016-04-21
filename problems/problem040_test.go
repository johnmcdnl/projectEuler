package problems

import "testing"

func BenchmarkProblem040(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem040()
	}
}

