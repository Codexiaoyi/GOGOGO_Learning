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

//TODO
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
	//用一个变量保存向前的节点指针
	var preNode *ListNode
	for head != nil {
		//首先临时指针向后一个节点迭代
		temp := head.Next
		//当前节点指向前一个节点
		head.Next = preNode
		//向前指针指向当前，作为下一个节点的前置节点
		preNode = head
		//当前指针向后移动
		head = temp
	}
	return preNode
}

//********************************92. 反转链表 II*********************************
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummyNode := &ListNode{Next: head}
	preNode, postNode := dummyNode, dummyNode
	var index int = 1
	//第一遍记录需要反转部分的前置节点和后置节点
	for index <= right {
		if index == left-1 {
			preNode = head
		}
		if index == right {
			postNode = head.Next
		}
		index++
		head = head.Next
	}

	head = preNode.Next
	curPreNode := postNode
	//第二遍反转需要反转的部分，再链接前置和后置节点
	for head != postNode {
		temp := head.Next
		head.Next = curPreNode
		curPreNode = head
		head = temp
	}
	preNode.Next = curPreNode

	return dummyNode.Next
}

//********************************21. 合并两个有序链表*********************************
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Next: l1}
	current := dummyNode
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			current.Next = l2
			break
		}
		if l1 != nil && l2 == nil {
			current.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	return dummyNode.Next
}

//TODO
//********************************86. 分隔链表*********************************
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Val: 0}

	return dummyNode.Next
}

//********************************148. 排序链表*********************************
//********************************143. 重排链表*********************************
//********************************141. 环形链表*********************************
//********************************142. 环形链表 II*********************************
//********************************234. 回文链表*********************************
//********************************138. 复制带随机指针的链表*********************************
