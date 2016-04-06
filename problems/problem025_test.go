package problems

import (
	"testing"
)

func BenchmarkProblem025(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem025()
	}
}
