package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem047(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem047()
	}
}

func TestProblem047(t *testing.T) {
	assert.Equal(t, Problem047(), 134043)
}