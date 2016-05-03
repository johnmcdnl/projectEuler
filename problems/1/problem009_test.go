package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem009(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem009()
	}
}

func TestProblem009(t *testing.T) {
	assert.Equal(t, Problem009(), 31875000)
}
