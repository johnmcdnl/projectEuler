package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem025(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem025()
	}
}

func TestProblem025(t *testing.T) {
	assert.Equal(t, Problem025(), int(4782))
}