package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem010(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem010()
	}
}
func TestProblem010(t *testing.T) {
	assert.Equal(t, 142913828922, Problem010())
}
