package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem001(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem001()
	}
}

func TestProblem001(t *testing.T) {
	assert.Equal(t, Problem001(), 233168)
}