package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem019(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem019()
	}
}

func TestProblem019(t *testing.T) {
	assert.Equal(t, Problem019(), 171)
}