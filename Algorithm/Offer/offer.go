package offer

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"
)

//****************************************剑指 Offer 04. 二维数组中的查找****************************************
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
// 		return false
// 	}
// 	row := len(matrix)
// 	col := len(matrix[0])
// 	for i := 0; i < row; i++ {
// 		if matrix[i][col-1] < target {
// 			continue
// 		}
// 		for j := 0; j < col; j++ {
// 			if matrix[i][j] == target {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

//****************************************剑指 Offer 05. 替换空格****************************************
// func replaceSpace(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	if len(s) == 1 {
// 		if s == " " {
// 			return "%20"
// 		} else {
// 			return s
// 		}
// 	}
// 	midle := len(s) / 2
// 	left := replaceSpace(s[:midle])
// 	right := replaceSpace(s[midle:])
// 	return left + right
// }

//****************************************剑指 Offer 09.用两个栈实现队列****************************************

type CQueue struct {
	stack_set    []int
	stack_delete []int
}

func Constructor_CQueue() CQueue {
	return CQueue{stack_set: make([]int, 0), stack_delete: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.stack_set = append(this.stack_set, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack_delete) == 0 {
		sl := len(this.stack_set)
		if sl == 0 {
			return -1
		}

		for i := 0; i < sl; i++ {
			if len(this.stack_set) == 0 {
				this.stack_delete = append(this.stack_delete, -1)
			} else {
				this.stack_delete = append(this.stack_delete, this.stack_set[len(this.stack_set)-1])
				this.stack_set = this.stack_set[:len(this.stack_set)-1]
			}
		}
	}

	res := this.stack_delete[len(this.stack_delete)-1]
	this.stack_delete = this.stack_delete[:len(this.stack_delete)-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

//****************************************剑指 Offer 30. 包含min函数的栈****************************************
type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	cur := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.min) != 0 && this.min[len(this.min)-1] == cur {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

//****************************************剑指 Offer 06. 从尾到头打印链表***************************************
func reversePrint(head *ListNode) []int {
	tmp := make([]int, 0)
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	res := make([]int, 0, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

//****************************************剑指 Offer 24. 反转链表***************************************
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//****************************************剑指 Offer 35. 复杂链表的复制***************************************
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	dummyNode := &Node{Next: head}
	dic := make(map[*Node]int)
	m := make(map[int]*Node) //index, node
	count := 0
	for head != nil {
		m[count] = &Node{Val: head.Val}
		dic[head] = count
		count++
		head = head.Next
	}
	head = dummyNode.Next
	count = 0
	for head != nil {
		m[count].Next = m[count+1]
		if head.Random != nil {
			index := dic[head.Random]
			m[count].Random = m[index]
		}
		count++
		head = head.Next
	}
	return m[0]
}

//****************************************剑指 Offer 05. 替换空格***************************************
func replaceSpace(s string) string {
	b := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

//****************************************剑指 Offer 58 - II. 左旋转字符串***************************************
func reverseLeftWords(s string, n int) string {
	b := strings.Builder{}
	b.WriteString(s[n:])
	b.WriteString(s[:n])
	return b.String()
}

//****************************************剑指 Offer 03. 数组中重复的数字***************************************
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return nums[i]
		}
		m[nums[i]] = 1
	}
	return -1
}

//****************************************剑指 Offer 53 - I. 在排序数组中查找数字 I***************************************
func search(nums []int, target int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	return m[target]
}

//****************************************剑指 Offer 53 - II. 0～n-1中缺失的数字***************************************
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

//****************************************剑指 Offer 04. 二维数组中的查找***************************************
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		if matrix[i][col-1] < target {
			continue
		}
		for j := 0; j < col; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//****************************************剑指 Offer 11. 旋转数组的最小数字***************************************
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

//****************************************剑指 Offer 50. 第一个只出现一次的字符***************************************
func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

//****************************************剑指 Offer 32 - I. 从上到下打印二叉树***************************************
func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	length := len(queue)
	for length != 0 {
		for i := 0; i < length; i++ {
			node := queue[0]
			res = append(res, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		length = len(queue)
	}
	return res
}

//****************************************剑指 Offer 32 - II. 从上到下打印二叉树 II***************************************
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	length := len(queue)
	for length != 0 {
		row := make([]int, length)
		for i := 0; i < length; i++ {
			node := queue[0]
			row[i] = node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, row)
		length = len(queue)
	}
	return res
}

//****************************************剑指 Offer 32 - III. 从上到下打印二叉树 III***************************************
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	length := len(queue)
	l2r := true
	for length != 0 {
		row := make([]int, length)
		if l2r {
			for i := 0; i < length; i++ {
				node := queue[0]
				row[i] = node.Val
				queue = queue[1:]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				node := queue[0]
				row[i] = node.Val
				queue = queue[1:]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		res = append(res, row)
		length = len(queue)
		l2r = !l2r
	}
	return res
}

//***************!!!!!!!!**************剑指 Offer 26. 树的子结构***************!!!!!!!!***************
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, A)
	l := len(queue)
	for l != 0 {
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Val == B.Val {
				if isSame(node, B) {
					return true
				}
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		l = len(queue)
	}
	return false
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	left := isSame(A.Left, B.Left)
	right := isSame(A.Right, B.Right)
	if left && right {
		return true
	}
	return false
}

//****************************************剑指 Offer 27. 二叉树的镜像***************************************
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var temp *TreeNode
	if root.Left != nil {
		temp = root.Left
	}
	root.Left = root.Right
	root.Right = temp
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

//****************************************剑指 Offer 28. 对称的二叉树***************************************
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && compare(a.Left, b.Right) && compare(a.Right, b.Left)
}

//****************************************剑指 Offer 10- I. 斐波那契数列***************************************
func fib(n int) int {
	if n < 2 {
		return n
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1]%(1e9+7) + f[i-2]%(1e9+7)
	}
	return f[n] % (1e9 + 7)
}

//****************************************剑指 Offer 10- II. 青蛙跳台阶问题***************************************
func numWays(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]%(1e9+7) + dp[i-2]%(1e9+7)
	}
	return dp[n] % (1e9 + 7)
}

//****************************************剑指 Offer 63. 股票的最大利润***************************************
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		res = int(math.Max(float64(res), float64(prices[i]-min)))
		min = int(math.Min(float64(min), float64(prices[i])))
	}
	return res
}

//****************************************剑指 Offer 42. 连续子数组的最大和**************************************
func maxSubArray(nums []int) int {
	max := nums[0]
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		new := nums[i]
		sum := nums[i] + f[i-1]
		if new > sum {
			f[i] = new
		} else {
			f[i] = sum
		}
		if f[i] > max {
			max = f[i]
		}
	}
	return max
}

//****************************************剑指 Offer 47. 礼物的最大价值**************************************
func maxValue(grid [][]int) int {
	row_len := len(grid)
	col_len := len(grid[0])
	f := make([][]int, row_len)
	for i := 0; i < row_len; i++ {
		f[i] = make([]int, col_len)
	}
	f[0][0] = grid[0][0]
	for i := 1; i < row_len; i++ {
		f[i][0] = grid[i][0] + f[i-1][0]
	}
	for i := 1; i < col_len; i++ {
		f[0][i] = grid[0][i] + f[0][i-1]
	}
	for i := 1; i < row_len; i++ {
		for j := 1; j < col_len; j++ {
			up := f[i][j-1]
			left := f[i-1][j]
			if up >= left {
				f[i][j] = up + grid[i][j]
			} else {
				f[i][j] = left + grid[i][j]
			}
		}
	}
	return f[row_len-1][col_len-1]
}

//****************************************剑指 Offer 46. 把数字翻译成字符串**************************************
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	nums := strconv.Itoa(num)
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		pre, _ := strconv.Atoi(string(nums[i-1]))
		cur, _ := strconv.Atoi(string(nums[i]))
		if pre == 0 || pre*10+cur > 25 {
			f[i] = f[i-1]
		} else {
			if i == 1 {
				f[i] = 2
			} else {
				f[i] = f[i-1] + f[i-2]
			}
		}
	}
	return f[len(nums)-1]
}

//****************************************剑指 Offer 48. 最长不含重复字符的子字符串**************************************
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	start, end := 0, 0
	maxLength := 0
	m := make(map[byte]bool)
	for start < len(s) {
		if start != 0 {
			delete(m, s[start-1])
		}
		for end < len(s) && !m[s[end]] {
			m[s[end]] = true
			end++
		}
		cur := end - start
		if cur > maxLength {
			maxLength = cur
		}
		start++
	}
	return maxLength
}

//****************************************剑指 Offer 18. 删除链表的节点**************************************
func deleteNode(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	cur := dummyNode
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return dummyNode.Next
}

//****************************************剑指 Offer 22. 链表中倒数第k个节点**************************************
func getKthFromEnd(head *ListNode, k int) *ListNode {
	count := 0
	slow, fast := head, head
	for fast != nil {
		if count == k {
			slow = slow.Next
		} else {
			count++
		}
		fast = fast.Next
	}
	return slow
}

//****************************************剑指 Offer 25. 合并两个排序的链表**************************************
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummyNode.Next
}

//****************************************剑指 Offer 52. 两个链表的第一个公共节点**************************************
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{})
	for {
		if headA == nil {
			break
		}
		nodeMap[headA] = struct{}{}
		headA = headA.Next
	}

	for {
		if headB == nil {
			return nil
		}
		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
}

//****************************************剑指 Offer 21. 调整数组顺序使奇数位于偶数前面**************************************
func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]%2 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

//****************************************剑指 Offer 57. 和为s的两个数字**************************************
func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	front, behind := 0, len(nums)-1
	for front != behind {
		if nums[front]+nums[behind] == target {
			res = append(res, nums[front], nums[behind])
			break
		}
		if nums[front]+nums[behind] > target {
			behind--
		}
		if nums[front]+nums[behind] < target {
			front++
		}
	}
	return res
}

//****************************************剑指 Offer 58 - I. 翻转单词顺序**************************************
func reverseWords(s string) string {
	var res string
	var i = len(s) - 1
	var j = i
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res += s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

//****************************************剑指 Offer 12. 矩阵中的路径**************************************

//****************************************剑指 Offer 45. 把数组排成最小的数**************************************
// func minNumber(nums []int) string {
// 	h := (IHeap)(nums)
// 	heap.Init(&h)
// 	var b strings.Builder
// 	for i := 0; i < h.Len(); i++ {
// 		b.WriteString(strconv.Itoa(h[i]))
// 	}
// 	return b.String()
// }

// type IHeap []int

// func (h IHeap) Len() int {
// 	return len(h)
// }

// func (h IHeap) Less(i, j int) bool {
// 	front := strconv.Itoa(h[i])
// 	back := strconv.Itoa(h[j])
// 	index := 0
// 	for index < len(front) && index < len(back) {
// 		f, _ := strconv.Atoi(string(front[index]))
// 		b, _ := strconv.Atoi(string(back[index]))
// 		if f < b {
// 			return true
// 		}
// 		if f > b {
// 			return false
// 		}
// 		index++
// 	}
// 	return index < len(front)
// }

// func (h IHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *IHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *IHeap) Pop() interface{} {
// 	n := len(*h)
// 	x := (*h)[n-1]
// 	*h = (*h)[:n-1]
// 	return x
// }

//****************************************剑指 Offer 61. 扑克牌中的顺子**************************************
func isStraight(nums []int) bool {
	sort.Ints(nums)
	i := 0
	for nums[i] == 0 {
		i++
	}
	if i > 1 {
		return false
	}
	if nums[4]-nums[i] != 4 {
		return false
	}
	for j := i + 1; j < 5; j++ {
		if nums[j] == nums[j-1] {
			return false
		}
	}
	return true
}

//****************************************剑指 Offer 40. 最小的k个数**************************************
func getLeastNumbers(arr []int, k int) []int {
	h := heap40{}
	heap.Init(&h)
	for i := 0; i < len(arr); i++ {
		heap.Push(&h, arr[i])
	}
	for i := 0; i < len(arr)-k; i++ {
		heap.Pop(&h)
	}
	return h
}

type heap40 []int

func (h heap40) Len() int {
	return len(h)
}

func (h heap40) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h heap40) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heap40) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heap40) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}

//****************************************剑指 Offer 55 - I. 二叉树的深度**************************************
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

//****************************************剑指 Offer 55 - II. 平衡二叉树**************************************
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//分治法
	balanced, _ := childBalanced(root)
	return balanced
}

func childBalanced(node *TreeNode) (isBalanced bool, floor float64) {
	if node == nil {
		return true, 0
	}

	lb, lf := childBalanced(node.Left)
	rb, rf := childBalanced(node.Right)
	if !lb || !rb {
		return false, 0
	}

	floor = math.Abs(float64(lf - rf))
	if floor > 1 {
		return false, floor
	}
	if lf > rf {
		return true, lf + 1
	}
	return true, rf + 1
}

//****************************************剑指 Offer 64. 求1+2+…+n**************************************
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

//****************************************剑指 Offer 68 - II. 二叉树的最近公共祖先**************************************
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor, _ := lowestCommonAncestor_dfs(root, p, q)
	return ancestor
}

func lowestCommonAncestor_dfs(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root == p || root == q {
		return root, true
	}
	left, le := lowestCommonAncestor_dfs(root.Left, p, q)
	right, re := lowestCommonAncestor_dfs(root.Right, p, q)
	if !le && !re {
		return nil, false
	}
	if !le && re {
		return right, true
	}
	if !re && le {
		return left, true
	}
	return root, true
}

//****************************************剑指 Offer 07. 重建二叉树**************************************
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	index := 0
	for index < len(inorder) {
		if inorder[index] == preorder[0] {
			break
		}
		index++
	}
	//[3,9,10, 20,15,7] 2
	//[9,10,3, 15,20,7]
	left := buildTree(preorder[1:index+1], inorder[:index])
	right := buildTree(preorder[index+1:], inorder[index+1:])
	return &TreeNode{Val: preorder[0], Left: left, Right: right}
}

//****************************************剑指 Offer 33. 二叉搜索树的后序遍历序列**************************************
// func verifyPostorder(postorder []int) bool {
// 	if len(postorder) == 0 {
// 		return true
// 	}
// 	_, _, res := verifyPostorder_for(postorder)
// 	return res
// }

// func verifyPostorder_for(postorder []int) (int, int, bool) {
// 	if len(postorder) == 0 {
// 		return math.MaxInt, math.MinInt, true
// 	}
// 	if len(postorder) == 1 {
// 		return postorder[0], postorder[0], true
// 	}
// 	cur := postorder[len(postorder)-1]
// 	index := len(postorder) - 2
// 	for index >= 0 {
// 		if postorder[index] < cur {
// 			break
// 		}
// 		index--
// 	}
// 	lmax, _, left := verifyPostorder_for(postorder[:index+1])
// 	_, rmin, right := verifyPostorder_for(postorder[index+1 : len(postorder)-1])
// 	if left && right && (lmax == math.MaxInt || lmax < cur) && (rmin == math.MinInt || rmin > cur) {
// 		return postorder[len(postorder)-2], postorder[0], true
// 	}
// 	return math.MaxInt, math.MinInt, false
// }

//*************************************剑指 Offer 15. 二进制中1的个数**************************************
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}
