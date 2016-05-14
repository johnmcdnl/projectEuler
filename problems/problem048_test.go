package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem048(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem048()
	}
}

func TestProblem048(t *testing.T) {
	assert.Equal(t, 9110846700, Problem048())
}