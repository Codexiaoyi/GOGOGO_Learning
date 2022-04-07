package Sort

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	example := []int{9, 3, 5, 22, 6, 4, 3}
	sortQuick(example, 0, len(example)-1)
	assert.Equal(t, []int{3, 3, 4, 5, 6, 9, 22}, example)
}

func TestFindKthLargest(t *testing.T) {
	example := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	res := findKthLargest(example, 20)
	assert.Equal(t, 2, res)
}

func TestTopKFrequent(t *testing.T) {
	example := []int{4, 1, -1, 2, -1, 2, 3}
	res := topKFrequent(example, 2)
	assert.Equal(t, []int{-1, 2}, res)
}

func TestDown(t *testing.T) {
	example := []int{4, 1, -1, 2, -1, 2, 3}
	down(math.MaxInt, math.MaxInt)
	assert.Equal(t, []int{-1, 2}, example)
}

func down(i0, n int) bool {
	i := i0
	for {
		j1 := i * 3
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n {
			j = j2 // = 2*i + 2  // right child
		}
		i = j
	}
	return i > i0
}
