package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkProblem022(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Problem022()
	}
}

/*
For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
*/
func TestProblem022(t *testing.T) {
	assert.Equal(t, 871198282, Problem022())
}

func TestSpecificNameValue(t *testing.T) {
	assert.Equal(t, nameValue("COLIN"), 53)
}

func TestDataOrdering(t *testing.T) {
	assert.Equal(t, sortedDataset()[937], "COLIN")
}
