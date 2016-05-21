package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem017(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem017()
	}
}
func TestProblem017(t *testing.T) {
	assert.Equal(t, 21124, Problem017())
}
