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
func sortList1(head *ListNode) *ListNode {
	//超时
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{Next: head}
	current := dummyNode
	var isSort bool = true
	for isSort {
		//是否有排序
		isSort = false
		for current.Next != nil && current.Next.Next != nil {
			if current.Next.Val > current.Next.Next.Val {
				//交换
				temp0 := current.Next.Next
				temp1 := current.Next.Next.Next
				current.Next.Next.Next = current.Next
				current.Next.Next = temp1
				current.Next = temp0
				//有发生排序
				isSort = true
			}
			current = current.Next
		}
		//走完一遍就从头来
		current = dummyNode
	}
	return dummyNode.Next
}

func sortList(head *ListNode) *ListNode {
	//本题目要求要用O(nlogn)时间和O(1)空间
	if head == nil || head.Next == nil {
		return head
	}
	//归并排序思想
	dummyNode := &ListNode{Val: 0}
	newHead := dummyNode
	slow, fast := head, head.Next
	//快指针走两步，慢指针走一步，快指针走完，慢指针就会停在一半的位置
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//**需要断开中间节点**
	middle := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(middle)
	for left != nil || right != nil {
		if left == nil {
			//只有右边有
			newHead.Next = right
			break
		}
		if right == nil {
			//只有左边有
			newHead.Next = left
			break
		}
		if left.Val <= right.Val {
			newHead.Next = left
			left = left.Next
		} else {
			newHead.Next = right
			right = right.Next
		}
		newHead = newHead.Next
	}
	return dummyNode.Next
}

//********************************143. 重排链表*********************************
//看错题目版本，这个是隔一个拿出来重新排列
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//思路：1.分类 2.反转 3.组合
	cur := head
	//分类
	reverseList := &ListNode{Val: 0}
	newHead := reverseList
	for cur != nil && cur.Next != nil {
		newHead.Next = cur.Next
		cur.Next = cur.Next.Next
		newHead.Next.Next = nil
		newHead = newHead.Next
		cur = cur.Next
	}
	//反转新的链表
	cur = reverseList.Next
	var preNode *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = temp
	}
	cur = head
	//组合
	for preNode != nil {
		temp := preNode.Next
		preNode.Next = cur.Next
		cur.Next = preNode
		preNode = temp
		cur = cur.Next.Next
	}
}

//这个是题目要求的，ps：这个解法效率不够
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//思路：1.找中间点 2.切割 3.反转 4.组合
	//找中间点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//切割
	middle := slow.Next
	slow.Next = nil

	//反转
	var preNode *ListNode
	for middle != nil {
		temp := middle.Next
		middle.Next = preNode
		preNode = middle
		middle = temp
	}

	cur := head
	//组合
	for preNode != nil {
		temp := preNode.Next
		preNode.Next = cur.Next
		cur.Next = preNode
		preNode = temp
		cur = cur.Next.Next
	}
}

//********************************141. 环形链表*********************************
func hasCycle(head *ListNode) bool {
	//环的问题离不开快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//********************************142. 环形链表 II*********************************
func detectCycle(head *ListNode) *ListNode {
	//依旧快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			//有环
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

//********************************234. 回文链表*********************************
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//找中间
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	middle := slow.Next
	slow.Next = nil

	var preNode *ListNode
	for middle != nil {
		temp := middle.Next
		middle.Next = preNode
		preNode = middle
		middle = temp
	}
	for head != nil && preNode != nil {
		if head.Val != preNode.Val {
			return false
		}
		head = head.Next
		preNode = preNode.Next
	}
	return true
}

//****************************************707. 设计链表****************************************
type MyLinkedList struct {
	Length     int
	Head, Tail *LinkNode
}
type LinkNode struct {
	Val        int
	Prev, Next *LinkNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Length {
		return -1
	}
	return this.getNode(index).Val
}

func (this *MyLinkedList) getNode(index int) *LinkNode {
	if index <= this.Length/2 {
		//前半段
		temp := this.Head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
		return temp
	} else {
		temp := this.Tail
		for i := 0; i < this.Length-index-1; i++ {
			temp = temp.Prev
		}
		return temp
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head != nil {
		//有头节点
		temp := &LinkNode{Val: val, Next: this.Head}
		this.Head.Prev = temp
		//维护头节点
		this.Head = temp
	} else {
		//新加头节点
		this.Head = &LinkNode{Val: val}
		this.Tail = this.Head
	}
	this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Tail != nil {
		//有尾部节点
		temp := &LinkNode{Val: val, Prev: this.Tail}
		this.Tail.Next = temp
		this.Tail = temp
	} else {
		this.Head = &LinkNode{Val: val}
		this.Tail = this.Head
	}
	this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if this.Length == index {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	node := this.getNode(index)
	temp := &LinkNode{Val: val}
	node.Prev.Next = temp
	temp.Prev = node.Prev
	node.Prev = temp
	temp.Next = node
	this.Length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < this.Length {
		node := this.getNode(index)
		if node == this.Head {
			if this.Length == 1 {
				this.Head = nil
			} else {
				this.Head = node.Next
				this.Head.Prev = nil
				node.Next = nil
			}
		} else if node == this.Tail {
			this.Tail = node.Prev
			this.Tail.Next = nil
			node.Prev = nil
		} else {
			//中间节点
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			node.Prev = nil
			node.Next = nil
		}
		this.Length--
	}
}
