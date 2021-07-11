package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test03(t *testing.T) {
	assert.Equal(t, 1, findRepeatNumber([]int{1, 1, 1}))
}
