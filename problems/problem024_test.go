package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem024(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem024()
	}
}
func TestProblem024(t *testing.T) {
	assert.Equal(t, nil, Problem024())
}
