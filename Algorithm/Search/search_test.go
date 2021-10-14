package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{2, 3, 5, 7, 8, 22}
	res1 := binarySearch(nums, 2)
	res2 := binarySearch(nums, 4)
	assert.True(t, res1)
	assert.False(t, res2)
}
