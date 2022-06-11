package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test03(t *testing.T) {
	assert.Equal(t, false, findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	assert.Equal(t, true, findNumberIn2DArray([][]int{{-1, 3}}, -1))
	assert.Equal(t, true, findNumberIn2DArray([][]int{{-1}, {3}}, -1))
}
