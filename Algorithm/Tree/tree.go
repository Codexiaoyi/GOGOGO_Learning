package tree

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//****************************************前序遍历****************************************
func preTravelsal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

//****************************************中序遍历****************************************
func inTravelsal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			//全部放入栈中，往左走
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

//****************************************后序遍历****************************************
func postTravelsal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	lastNode := new(TreeNode)
	for root != nil || len(stack) != 0 {
		for root != nil {
			//全部放入栈中，往左走
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		//如果没有右节点或者右节点已经被遍历，那么就弹出节点并且记录遍历的值，否则就往右走
		if root.Right == nil || root.Right == lastNode {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			lastNode = node
		} else {
			root = node.Right
		}
	}
	return result
}

//****************************************BFS层次遍历，一行一行遍历****************************************
func bfsTravelsal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	//队列表示当前层元素
	queue := make([]*TreeNode, 0)
	//加入树根元素
	queue = append(queue, root)
	//一层都没元素就结束
	for len(queue) != 0 {
		//计算当前层有几个元素（都在队列里）
		l := len(queue)
		//遍历当前层元素
		for i := 0; i < l; i++ {
			//先将当前节点记录，退出队列
			node := queue[0]
			result = append(result, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

//****************************************104二叉树的最大深度****************************************
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

//****************************************110平衡二叉树****************************************
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

//****************************************236二叉树的最近公共祖先****************************************
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//分治法，如果两个节点在左子树有公共祖先或者右子树有，那就直接返回左子树或右子树，没有则返回当前根节点。
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//左右两边如果都不为空，则根节点就是公共祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

//****************************************102二叉树的层序遍历****************************************
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	//使用队列保存每层的节点
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	l := len(queue)
	for l > 0 {
		//只要当前层还有元素就要遍历
		level := make([]int, 0)
		for i := 0; i < l; i++ {
			//每一层的操作
			node := queue[0]
			queue = queue[1:] //出队
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//下一层的节点个数
		l = len(queue)
		result = append(result, level)
	}

	return result
}

//****************************************107二叉树的层序遍历二（自底向上层序遍历）****************************************
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	//使用栈保存正序
	stack := make([][]int, 0)
	//使用队列保存每层的节点
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	l := len(queue)
	for l > 0 {
		//只要当前层还有元素就要遍历
		level := make([]int, 0)
		for i := 0; i < l; i++ {
			//每一层的操作
			node := queue[0]
			queue = queue[1:] //出队
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//下一层的节点个数
		l = len(queue)
		stack = append(stack, level)
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append(result, node)
		//弹出
		stack = stack[:len(stack)-1]
	}

	return result
}

//****************************************103二叉树的锯齿形层序遍历（一行从左到右，一行从右到左）****************************************
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false
	for len(queue) > 0 {
		list := make([]int, 0)
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		if toggle {
			for i := 0; i < len(list)/2; i++ {
				list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
			}
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}

//**************************************** 98验证二叉搜索树 **************************************
type ResultType struct {
	IsValid bool
	// 记录左右两边最大最小值，和根节点进行比较
	Max *TreeNode
	Min *TreeNode
}

func isValidBST(root *TreeNode) bool {
	result := helper(root)
	return result.IsValid
}
func helper(root *TreeNode) ResultType {
	result := ResultType{}
	// check
	if root == nil {
		result.IsValid = true
		return result
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}
	if left.Max != nil && left.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}

	result.IsValid = true
	// 如果左边还有更小的3，就用更小的节点，不用4
	//  5
	// / \
	// 1   4
	//      / \
	//     3   6
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}
	return result
}

//****************************************701二叉搜索树中的插入操作（分治法做法）****************************************
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	//利用分治法直接插到相应叶子节点下
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//****************************************100. 相同的树****************************************
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		left := isSameTree(p.Left, q.Left)
		right := isSameTree(p.Right, q.Right)

		if left && right {
			return true
		}
	}

	return false
}

//****************************************226. 翻转二叉树****************************************
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	t := root.Left
	root.Left = root.Right
	root.Right = t

	return root
}

//****************************************257. 二叉树的所有路径****************************************
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}

	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	if len(left) == 0 && len(right) == 0 {
		res = append(res, strconv.Itoa(root.Val))
	}

	for _, l := range left {
		res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(root.Val), l))
	}

	for _, r := range right {
		res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(root.Val), r))
	}
	return res
}

//****************************************101. 对称二叉树****************************************
//抄题解
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//左子树的左子树与右子树的右子树比，左子树的右子树与右子树的左子树比
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

//****************************************108. 将有序数组转换为二叉搜索树****************************************
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	//二叉搜索树中序遍历就是升序
	midle := len(nums) / 2
	root := &TreeNode{Val: nums[midle]}
	root.Left = sortedArrayToBST(nums[:midle])
	root.Right = sortedArrayToBST(nums[midle+1:])

	return root
}

//****************************************235. 二叉搜索树的最近公共祖先****************************************
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == q || root == p {
		return root
	}

	left := lowestCommonAncestor235(root.Left, p, q)
	right := lowestCommonAncestor235(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}

//****************************************105. 从前序与中序遍历序列构造二叉树****************************************
func buildTree(preorder []int, inorder []int) *TreeNode {
	//前序遍历的第一个是根节点,中序遍历根节点的左边都是左子树，右边是右子树
	if len(preorder) == 0 {
		return nil
	}
	//左子树中序遍历slice
	left_inorder := make([]int, 0)
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
		left_inorder = append(left_inorder, inorder[i])
	}
	//分割出左右子树的前序和中序
	right_inorder := inorder[len(left_inorder)+1:]
	right_preorder := preorder[len(left_inorder)+1:]
	left_preorder := preorder[1 : len(left_inorder)+1]
	//构造左右子树
	left := buildTree(left_preorder, left_inorder)
	right := buildTree(right_preorder, right_inorder)
	root := &TreeNode{Val: preorder[0], Left: left, Right: right}
	return root
}

//****************************************111. 二叉树的最小深度****************************************
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	//叶子节点
	if left == 0 && right == 0 {
		return 1
	}

	if left != 0 && right != 0 {
		if left >= right {
			return right + 1
		} else {
			return left + 1
		}
	}

	if left == 0 {
		return right + 1
	} else {
		return left + 1
	}
}

//****************************************404. 左叶子之和****************************************
func sumOfLeftLeaves(root *TreeNode) int {
	return travelsumOfLeftLeaves(root, false)
}

func travelsumOfLeftLeaves(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	//叶子
	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		} else {
			return 0
		}
	}

	left := travelsumOfLeftLeaves(root.Left, true)
	right := travelsumOfLeftLeaves(root.Right, false)

	return left + right
}

//****************************************112. 路径总和****************************************
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	left := hasPathSum(root.Left, targetSum-root.Val)
	right := hasPathSum(root.Right, targetSum-root.Val)

	return left || right
}

//****************************************199. 二叉树的右视图****************************************
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	//思路：层序遍历每层最后一个节点
	queue := make([]*TreeNode, 0)
	//根节点入队
	queue = append(queue, root)
	length := len(queue)
	for length != 0 {
		//每一层操作
		for i := 0; i < length; i++ {
			if i == length-1 {
				//最后一个节点
				res = append(res, queue[i].Val)
			}
			//左右节点入队
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		//一层结束
		queue = queue[length:]
		length = len(queue)
	}
	return res
}

//****************************************129. 求根节点到叶节点数字之和****************************************
func sumNumbers(root *TreeNode) int {
	return travel_sumNumbers(root, 0)
}

func travel_sumNumbers(root *TreeNode, total int) int {
	if root == nil {
		return 0
	}
	//叶子节点
	if root.Left == nil && root.Right == nil {
		//返回
		return total*10 + root.Val
	}
	//非叶子节点，就计算后传下去
	return travel_sumNumbers(root.Left, total*10+root.Val) + travel_sumNumbers(root.Right, total*10+root.Val)
}

//****************************************114. 二叉树展开为链表****************************************
func flatten(root *TreeNode) {
	//前序遍历后依次排入
	dummyNode := root
	var lastNode *TreeNode
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			lastNode = root
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//node节点放到最后一个节点左边
		lastNode.Left = node.Right
		root = node.Right
	}

	root = dummyNode
	for root != nil {
		root.Right = root.Left
		root.Left = nil
		root = root.Right
	}

	root = dummyNode
}

//****************************************230. 二叉搜索树中第K小的元素****************************************
func kthSmallest(root *TreeNode, k int) int {
	//中序遍历有序然后再找
	flag := 0
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if flag == k-1 {
			return node.Val
		}
		flag++
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return 0
}

//****************************************637. 二叉树的层平均值****************************************
func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		total := 0
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			total += node.Val
		}
		result = append(result, float64(total)/float64(length))
		queue = queue[length:]
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//****************************************109. 有序链表转换二叉搜索树****************************************
func sortedListToBST(head *ListNode) *TreeNode {
	nodes := make([]int, 0)
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	return sortedListToBST_buildTree(nodes)
}

func sortedListToBST_buildTree(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	left := sortedListToBST_buildTree(nodes[:mid])
	right := sortedListToBST_buildTree(nodes[mid+1:])
	return &TreeNode{Val: nodes[mid], Left: left, Right: right}
}

type Node struct {
	Val      int
	Children []*Node
	Left     *Node
	Right    *Node
	Next     *Node
}

//****************************************559. N 叉树的最大深度****************************************
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.Children {
		cd := maxDepth1(child)
		if cd > max {
			max = cd
		}
	}
	return max + 1
}

//****************************************700. 二叉搜索树中的搜索****************************************
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}
