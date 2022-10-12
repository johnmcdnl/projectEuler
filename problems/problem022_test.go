package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkProblem022(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem022()
	}
}
func TestProblem022(t *testing.T) {
	assert.Equal(t, 4179871, Problem022())
}
