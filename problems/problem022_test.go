package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem022(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem022()
	}
}
func TestProblem022(t *testing.T) {
	assert.Equal(t, nil, Problem022())
}
