package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem028(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem028()
	}
}
func TestProblem028(t *testing.T) {
	assert.Equal(t, 669171001, Problem028())
}
