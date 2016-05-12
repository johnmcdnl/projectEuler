package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem005(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem005()
	}
}

func TestProblem005(t *testing.T) {
	assert.Equal(t, 2520, Problem005())
}