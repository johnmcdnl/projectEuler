package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem067(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem067()
	}
}

func TestProblem067(t *testing.T) {
	assert.Equal(t, 7273, Problem067())
}