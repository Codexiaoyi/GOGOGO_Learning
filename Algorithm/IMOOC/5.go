package imooc

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表
//206
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

//92
//83
