package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem037(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem037()
	}
}

func TestProblem037(t *testing.T) {
	assert.Equal(t, 748317, Problem037())
}