package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem007(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem007()
	}
}

func TestProblem007(t *testing.T) {
	assert.Equal(t, Problem007(), 104743)
}