package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem049(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem049()
	}
}

func TestProblem049(t *testing.T) {
	assert.Equal(t, nil, Problem049())
}