package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem043(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem043()
	}
}

func TestProblem043(t *testing.T) {
	assert.Equal(t, Problem043(), 16695334890)
}