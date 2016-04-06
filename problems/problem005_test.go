package problems

import (
	"testing"
)

func BenchmarkProblem005(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem005()
	}
}
