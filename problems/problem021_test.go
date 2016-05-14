package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem021(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem021()
	}
}
func TestProblem021(t *testing.T) {
	assert.Equal(t, 31626, Problem021())
}

