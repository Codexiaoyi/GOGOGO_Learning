package dp

import (
	"math"
	"strconv"
)

/*
动态规划:
1.使用场景：如果一个问题，可以把所有可能的答案穷举出来，并且穷举出来后，发现存在重叠子问题，就可以考虑使用动态规划
2.解题思路：
	核心思想就是拆分子问题，记住过往，减少重复计算。
	1）穷举分析
	2）确定边界,也就是dp数组初始化
	3）找出规律，确定最优子结构
	4）写出状态转移方程，也就是递推公式
*/

//****************************************509. 斐波那契数****************************************
func fib(n int) int {
	dp := make([]int, 31)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//****************************************53. 最大子序和****************************************
func maxSubArray(nums []int) int {
	max := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

//****************************************64. 最小路径和****************************************
func minPathSum(grid [][]int) int {
	//dp[i][j]表示第i行j列上的最小路径和
	if grid == nil || grid[0] == nil {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			min := int(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j])))
			dp[i][j] = min + grid[i][j]
		}
	}
	return dp[rows-1][cols-1]
}

//****************************************746. 使用最小花费爬楼梯****************************************
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		if dp[i-1] > dp[i-2] {
			dp[i] = dp[i-2] + cost[i]
		} else {
			dp[i] = dp[i-1] + cost[i]
		}
	}
	if dp[len(cost)-1] > dp[len(cost)-2] {
		return dp[len(cost)-2]
	}
	return dp[len(cost)-1]
}

//****************************************213. 打家劫舍 II****************************************
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	fromHead := make([]int, len(nums))
	dp[0] = nums[0]
	fromHead[0] = 1
	if nums[0] > nums[1] {
		dp[1] = nums[0]
		fromHead[1] = 1
	} else {
		dp[1] = nums[1]
		fromHead[1] = 0
	}
	for i := 2; i < len(nums); i++ {
		if dp[i-1] >= dp[i-2]+nums[i] || (i == len(nums)-1 && fromHead[i-2] == 1) {
			dp[i] = dp[i-1]
			fromHead[i] = fromHead[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i]
			fromHead[i] = fromHead[i-2]
		}
	}
	return dp[len(nums)-1]
}

//****************************************1137. 第 N 个泰波那契数****************************************
func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

//****************************************70. 爬楼梯****************************************
func climbStairs(n int) int {
	if n >= 0 && n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//****************************************198. 打家劫舍****************************************
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+2)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

//****************************************120. 三角形最小路径和****************************************
// 记忆化搜索递归方法超时
// func minimumTotal(triangle [][]int) int {
// 	memorySearch := make([][]int, len(triangle))
// 	for i := 0; i < len(triangle); i++ {
// 		memorySearch[i] = make([]int, len(triangle[i]))
// 	}
// 	return minimumTotal_recursive(triangle, memorySearch, 0, 0)
// }

// func minimumTotal_recursive(triangle, memorySearch [][]int, curIndex_row, curIndex_col int) int {
// 	if curIndex_row == len(triangle) {
// 		return 0
// 	}
// 	if memorySearch[curIndex_row][curIndex_col] == 0 {
// 		left := minimumTotal_recursive(triangle, memorySearch, curIndex_row+1, curIndex_col)
// 		right := minimumTotal_recursive(triangle, memorySearch, curIndex_row+1, curIndex_col+1)
// 		if left >= right {
// 			return right + triangle[curIndex_row][curIndex_col]
// 		}
// 		return left + triangle[curIndex_row][curIndex_col]
// 	}
// 	return memorySearch[curIndex_row][curIndex_col]
// }
func minimumTotal(triangle [][]int) int {
	//dp[i][j]表示当前节点作为终点的最小路径和
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			if dp[i-1][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
		dp[i][len(triangle[i])-1] = dp[i-1][len(triangle[i-1])-1] + triangle[i][len(triangle[i])-1]
	}
	min := math.MaxInt32
	for i := 0; i < len(dp[len(dp)-1]); i++ {
		if dp[len(dp)-1][i] < min {
			min = dp[len(dp)-1][i]
		}
	}
	return min
}

//****************************************91. 解码方法****************************************
func numDecodings(s string) int {
	dp := make([]int, len(s))

	isValid := func(s string) bool {
		if s[0] == '0' {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if i >= 1 && i <= 26 {
			return true
		}
		return false
	}

	if isValid(string(s[0])) {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	if len(s) > 1 {
		if isValid(string(s[1])) {
			dp[1] += dp[0]
		}
		if isValid(string(s[:2])) {
			dp[1] += 1
		}
	}
	for i := 2; i < len(s); i++ {
		if isValid(string(s[i])) {
			dp[i] += dp[i-1]
		}
		if isValid(string(s[i-1 : i+1])) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)-1]
}

//****************************************5. 最长回文子串****************************************
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 根据case 1 初始数据
	}
	for length := 2; length <= len(s); length++ { //长度固定，不断移动起点
		for start := 0; start < len(s)-length+1; start++ { //长度固定，不断移动起点
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				dp[start][end] = true //即case 2的判断
			} else {
				dp[start][end] = dp[start+1][end-1] //状态转移
			}
			if dp[start][end] && (end-start+1) > len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}
	return result
}
