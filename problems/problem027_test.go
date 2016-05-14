package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem027(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem027()
	}
}
func TestProblem027(t *testing.T) {
	assert.Equal(t, nil, Problem027())
}
