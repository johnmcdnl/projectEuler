package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem040(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem040()
	}
}

func TestProblem040(t *testing.T) {
	assert.Equal(t, Problem040(), int(210))
}