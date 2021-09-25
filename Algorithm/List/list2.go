package List

//********************************83. 删除排序链表中的重复元素*********************************
func d83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Next: head}
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next, head.Next.Next = head.Next.Next, nil
		} else {
			head = head.Next
		}
	}
	return dummyNode.Next
}

//********************************82. 删除排序链表中的重复元素 II*********************************
