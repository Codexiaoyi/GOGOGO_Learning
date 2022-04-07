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

//********************************203*********************************
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	font := head
	back := head.Next
	for back != nil {
		if back.Val == val {
			font.Next = back.Next
			back.Next = nil
			back = font.Next
		} else {
			font = font.Next
			back = back.Next
		}
	}
	if head.Val == val {
		return head.Next
	}
	return head
}

//********************************86*********************************
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Next: head}
	partitionNode := dummyNode
	if head.Val < x {
		for head.Next != nil {
			//找大于等于x的前一个结点
			if head.Next.Val >= x {
				partitionNode = head
				break
			}
			head = head.Next
		}
		//所有都比x小
		if partitionNode == dummyNode {
			return dummyNode.Next
		}
	}

	temp := partitionNode.Next
	//从分割点后一个节点开始，小于x的都顺延放到分割点后
	for temp.Next != nil {
		if temp.Next.Val < x {
			lessNode := temp.Next
			temp.Next = temp.Next.Next
			lessNode.Next = partitionNode.Next
			partitionNode.Next = lessNode
			partitionNode = partitionNode.Next
		} else {
			temp = temp.Next
		}
	}
	return dummyNode.Next
}

//***********************147. 对链表进行插入排序*************************
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	pre, curr := head, head.Next
	for curr != nil {
		if pre.Val <= curr.Val {
			pre = pre.Next
		} else {
			pos := dummyHead
			for pos.Next.Val <= curr.Val {
				pos = pos.Next
			}
			pre.Next = curr.Next
			curr.Next = pos.Next
			pos.Next = curr
		}
		curr = pre.Next
	}
	return dummyHead.Next
}

//***********************328. 奇偶链表*************************
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	even := head.Next
	isOdd := true
	for fast != nil && fast.Next != nil {
		slow = fast
		fast = fast.Next
		slow.Next = fast.Next
		isOdd = !isOdd
	}
	if isOdd {
		//结尾是奇数
		fast.Next = even
	} else {
		slow.Next = even
	}
	return head
}

//***********************1669. 合并两个链表*************************
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	start, end := list1, list1
	list2Head := list2
	cur := list1
	count := 0
	for cur != nil {
		if count+1 == a {
			start = cur
		}
		if count-1 == b {
			end = cur
		}
		cur = cur.Next
		count++
		if list2.Next != nil {
			list2 = list2.Next
		}
	}
	for list2.Next != nil {
		list2 = list2.Next
	}
	start.Next = list2Head
	list2.Next = end
	return list1
}
