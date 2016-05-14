package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem041(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem041()
	}
}
func TestProblem041(t *testing.T) {
	assert.Equal(t, nil, Problem041())
}