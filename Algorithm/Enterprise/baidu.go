package enterprise

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//******************************************206反转链表
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var preNode *ListNode
	for head != nil {
		temp := head.Next
		head.Next = preNode
		preNode = head
		head = temp
	}
	return preNode
}

//******************************************215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	//构建大顶堆
	for i := length/2 - 1; i >= 0; i-- {
		sink(nums, i, length)
	}

	//取出最大元素并放到数组最后，最后元素提上来开始下沉
	for i := length - 1; i >= 1; i-- {
		//这个i表示长度，取出一个元素之后，数组长度就减小
		swap(nums, 0, i)
		sink(nums, 0, i)
	}

	return nums[len(nums)-k]
}

//i 当前节点的索引
//length
func sink(nums []int, i int, length int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		//左右和自己值最大的索引
		index := i
		//左节点存在且大于当前就记录
		if left < length && nums[left] > nums[index] {
			index = left
		}
		if right < length && nums[right] > nums[index] {
			index = right
		}
		//如果根节点最大直接结束循环
		if index == i {
			break
		}
		//交换并继续下沉
		swap(nums, index, i)
		i = index
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

//******************************************141.环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
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

//******************************************104.二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	res := 1
	if left >= right {
		res += left
	} else {
		res += right
	}

	return res
}

//******************************************1.两数之和
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//******************************************94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}

//******************************************142. 环形链表 II
//这里是答案，用hash，快慢指针需要数学推导，相遇后双指针分别从相遇点和头节点同时出发走一步，最后相遇的是第一个入环点
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}
