package window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	l := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, l)
}
