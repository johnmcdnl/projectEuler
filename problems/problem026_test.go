package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem026(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem026()
	}
}
func TestProblem026(t *testing.T) {
	assert.Equal(t, nil, Problem026())
}
