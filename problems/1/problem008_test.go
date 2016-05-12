package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem008(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem008()
	}
}

func TestProblem008(t *testing.T) {
	assert.Equal(t, 23514624000, Problem008())
}