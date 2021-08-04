package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	q := deleteAndEarn(nums)
	fmt.Println(q)
}
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	// sum[v] 代表 序号为v的 和，也就是nums中所有的v的和
	sum := make([]int, maxVal+1)
	for _, v := range nums {
		sum[v] += v
	}
	return rob(sum)
}

func rob(nums []int) int {
	//dp[i] 代表 第i个 删除的时候的所有的和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
