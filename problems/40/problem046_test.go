package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem046(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem046()
	}
}

func TestProblem046(t *testing.T) {
	assert.Equal(t, 5777, Problem046())
}