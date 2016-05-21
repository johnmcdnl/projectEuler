package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem042(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem042()
	}
}

func TestProblem042(t *testing.T) {
	assert.Equal(t, 162, Problem042())
}