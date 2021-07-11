package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test09(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(5)
	obj.AppendTail(2)
	param_1 := obj.DeleteHead()
	param_2 := obj.DeleteHead()
	assert.Equal(t, 5, param_1)
	assert.Equal(t, 2, param_2)
}
