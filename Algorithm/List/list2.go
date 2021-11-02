package List

//********************************83. 删除排序链表中的重复元素*********************************
func d83(head *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	if head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummyNode
}

//********************************206*********************************
