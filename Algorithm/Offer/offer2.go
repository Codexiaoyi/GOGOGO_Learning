package offer

import (
	"math"
	"sort"
)

//****************************************剑指 Offer II 007. 数组中和为 0 的三个数**************************************
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

//*****剑指 Offer II 008. 和大于等于 target 的最短子数组********
func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 1
	length := math.MaxInt
	sum := nums[0]
	for start < end {
		if sum >= target {
			if end-start < length {
				length = end - start
			}
			sum -= nums[start]
			start++
		} else {
			if end < len(nums) {
				sum += nums[end]
				end++
			}
		}
	}
	return length
}

//*****剑指 Offer II 009. 乘积小于 K 的子数组********
// func numSubarrayProductLessThanK(nums []int, k int) int {

// }

//*****112. 路径总和********
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	l := hasPathSum(root.Left, targetSum-root.Val)
	r := hasPathSum(root.Right, targetSum-root.Val)
	return l || r
}

//*****230. 二叉搜索树中第K小的元素********
func kthSmallest(root *TreeNode, k int) int {
	//中序遍历
	stack := make([]*TreeNode, 0)
	cur := 0
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		cur++
		if cur == k {
			return node.Val
		}
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return 0
}

//*****20. 有效的括号********
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
