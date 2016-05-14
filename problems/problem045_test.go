package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem045(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem045()
	}
}

func TestProblem045(t *testing.T) {
	assert.Equal(t, nil, Problem045())
}