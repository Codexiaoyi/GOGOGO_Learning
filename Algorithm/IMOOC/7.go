package imooc

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树与递归
//递归步骤：1.递归终止条件 2.递归过程

//104
//111
//226
//100

//101
//222
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return left + right + 1
}

//110
//112
//404
//113
func pathSum113(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		res = append(res, []int{root.Val})
		return res
	}

	left := pathSum113(root.Left, targetSum-root.Val)
	right := pathSum113(root.Right, targetSum-root.Val)

	for _, road := range left {
		newRoad := make([]int, 0, len(road)+1)
		newRoad = append(newRoad, root.Val)
		newRoad = append(newRoad, road...)
		res = append(res, newRoad)
	}

	for _, road := range right {
		newRoad := make([]int, 0, len(road)+1)
		newRoad = append(newRoad, root.Val)
		newRoad = append(newRoad, road...)
		res = append(res, newRoad)
	}

	return res
}

//129
//437
func pathSum(root *TreeNode, targetSum int) int {
	//递归遍历法
	if root == nil {
		return 0
	}

	res := includeRootPathSum(root, targetSum)
	res += pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
	return res

	//前缀和+

}
func dfs() {

}

func includeRootPathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == targetSum {
		res++
	}

	//包含当前节点
	res += includeRootPathSum(root.Left, targetSum-root.Val) + includeRootPathSum(root.Right, targetSum-root.Val)
	return res
}

//二分搜索树
//每个节点的键值都大于左孩子，每个节点的键值小于右孩子；以左右孩子为根的子树仍为二叉搜索树
//235
//98
//450
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Right == nil && root.Left != nil {
			return root.Left
		}
		//保存左右节点
		left := root.Left
		right := root.Right
		//temp查找右子树的最小节点
		temp := root.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		temp.Left = left
		return right
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

//108
//230
//236
