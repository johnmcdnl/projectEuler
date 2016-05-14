package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem044(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem044()
	}
}

func TestProblem044(t *testing.T) {
	assert.Equal(t, nil, Problem044())
}