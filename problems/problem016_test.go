package problems

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProblem016(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem016()
	}
}
func TestProblem016(t *testing.T) {
	assert.Equal(t, 1366, Problem016())
}
