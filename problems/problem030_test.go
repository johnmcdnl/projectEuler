package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem030(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem030()
	}
}
func TestProblem030(t *testing.T) {
	assert.Equal(t, 443839, Problem030())
}
