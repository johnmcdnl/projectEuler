package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem058(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem058()
	}
}

func TestProblem058(t *testing.T) {
	assert.Equal(t, 26241, Problem058())
}