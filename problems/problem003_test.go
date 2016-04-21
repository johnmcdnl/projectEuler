package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem003(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem003()
	}
}

func TestProblem003(t *testing.T) {
	assert.Equal(t, Problem003(), int(6857))
}