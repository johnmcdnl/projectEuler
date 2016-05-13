package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem006(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem006()
	}
}

func TestProblem006(t *testing.T) {
	assert.Equal(t, 25164150, Problem006())
}