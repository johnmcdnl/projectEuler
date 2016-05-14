package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem012(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem012()
	}
}
func TestProblem012(t *testing.T) {
	assert.Equal(t, 76576500, Problem012())
}
