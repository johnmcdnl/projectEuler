package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem002(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem002()
	}
}

func TestProblem002(t *testing.T) {
	assert.Equal(t, Problem002(), int(4613732))
}