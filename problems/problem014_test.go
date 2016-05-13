package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem014(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem014()
	}
}

func TestProblem014(t *testing.T) {
	assert.Equal(t, 837799, Problem014())
}