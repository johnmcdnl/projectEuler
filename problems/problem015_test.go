package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem015(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem015()
	}
}
func TestProblem015(t *testing.T) {
	assert.Equal(t, 137846528820, Problem015())
}
