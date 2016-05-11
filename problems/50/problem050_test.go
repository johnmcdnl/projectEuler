package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem050(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem050()
	}
}

func TestProblem050(t *testing.T) {
	assert.Equal(t, 997651, Problem050())
}