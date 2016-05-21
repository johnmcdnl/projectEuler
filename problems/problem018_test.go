package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem018(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem018()
	}
}
func TestProblem018(t *testing.T) {
	assert.Equal(t, 1074, Problem018())
}
