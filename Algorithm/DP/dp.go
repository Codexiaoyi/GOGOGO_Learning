package dp

import "math"

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
