package List

/*
链表的核心点：
1.null/nil 异常处理
2.dummy node 哑巴节点
3.快慢指针
4.插入一个节点到排序链表
5.从一个链表中移除一个节点
6.翻转链表
7.合并两个链表
8.找到链表的中间节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//********************************83. 删除排序链表中的重复元素*********************************
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummyNode.Next
}

//##############################################################################################
//********************************82. 删除排序链表中的重复元素 II*********************************
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Val: 0, Next: head}
	slow, fast := dummyNode, head
	for fast != nil {
		if fast.Next != nil && fast.Val != fast.Next.Val {
			if slow.Next == fast {
				slow = fast
			} else {
				slow.Next = fast.Next
				fast = slow
			}
		}
		if fast.Next == nil && slow.Next != fast {
			slow.Next = fast.Next
			fast = slow
		}
		fast = fast.Next
	}
	return dummyNode.Next
}

//********************************206. 反转链表*********************************
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

}

//********************************92. 反转链表 II*********************************
//********************************21. 合并两个有序链表*********************************
//********************************86. 分隔链表*********************************
//********************************148. 排序链表*********************************
//********************************143. 重排链表*********************************
//********************************141. 环形链表*********************************
//********************************142. 环形链表 II*********************************
//********************************234. 回文链表*********************************
//********************************138. 复制带随机指针的链表*********************************
