package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem023(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem023()
	}
}
func TestProblem023(t *testing.T) {
	assert.Equal(t, nil, Problem023())
}
