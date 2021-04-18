package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
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

//中序遍历
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

//后序遍历
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

//BFS层次遍历，一行一行遍历
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

//104二叉树的最大深度
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

//110平衡二叉树
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

//236二叉树的最近公共祖先
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

//102二叉树的层序遍历
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

//107二叉树的层序遍历二（自底向上层序遍历）
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

//103二叉树的锯齿形层序遍历（一行从左到右，一行从右到左）
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
