package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem013(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem013()
	}
}
func TestProblem013(t *testing.T) {
	assert.Equal(t, 5537376230, Problem013())
}
