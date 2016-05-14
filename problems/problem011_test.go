package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem011(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem011()
	}
}
func TestProblem011(t *testing.T) {
	assert.Equal(t, 70600674, Problem011())
}
