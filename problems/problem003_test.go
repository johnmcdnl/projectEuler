package problems

import "testing"

func BenchmarkProblem003(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem003()
	}
}
