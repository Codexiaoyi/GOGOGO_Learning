package List

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortList(t *testing.T) {
	root := BuildList([]int{3, 2, 4})
	newRoot := insertionSortList(root)
	actual := []int{}
	for newRoot != nil {
		actual = append(actual, newRoot.Val)
		newRoot = newRoot.Next
	}
	assert.Equal(t, []int{2, 3, 4}, actual)
}

func BuildList(nums []int) *ListNode {
	root := &ListNode{}
	cur := root
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return root.Next
}
