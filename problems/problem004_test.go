package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem004(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem004()
	}
}

func TestProblem004(t *testing.T) {
	assert.Equal(t, Problem004(), int(906609))
}