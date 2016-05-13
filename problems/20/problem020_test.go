package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem020(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem020()
	}
}
func TestProblem020(t *testing.T) {
	assert.Equal(t, 648, Problem020())
}
