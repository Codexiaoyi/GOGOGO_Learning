package Sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	example := []int{9, 3, 5, 22, 6, 4, 3}
	sortHeap(example)
	assert.Equal(t, []int{3, 3, 4, 5, 6, 9, 22}, example)
}
