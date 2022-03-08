package imooc

import "strconv"

//递归与回溯问题
//一般都可以做成树形结构，递归过程就是树形结构的纵向路线
//单次循环就是横向的节点数
//1.先确定终止条件，在终止条件中记录下单条路径到结果集
//2.开始横向的循环
//3.节点加入单条路径集（可能涉及剪枝）
//4.递归
//5.回溯一个单条路径的节点

//93
func restoreIpAddresses(s string) []string {
	res := []string{}
	ips := []string{}
	if len(s) < 4 {
		return res
	}
	restoreIpAddresses_back(s, 0, ips, &res)
	return res
}

func restoreIpAddresses_back(s string, index int, ips []string, res *[]string) {
	if len(s) == index && len(ips) == 4 {
		newRes := ips[0] + "." + ips[1] + "." + ips[2] + "." + ips[3]
		*res = append(*res, newRes)
	}

	for i := 1; i <= 3; i++ {
		if i+index > len(s) {
			return
		}
		ip := s[index : index+i]
		if !isIp(ip) {
			continue
		}
		ips = append(ips, ip)
		restoreIpAddresses_back(s, index+i, ips, res)
		ips = ips[:len(ips)-1]
	}
}

func isIp(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= 0 && i <= 255
}

//131
func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}
	memo := make(map[string]int)
	partition_back(s, 0, path, &res, memo)
	return res
}

func partition_back(s string, start int, path []string, res *[][]string, mem map[string]int) {
	if start == len(s) {
		t := make([]string, len(path))
		copy(t, path)
		*res = append(*res, t)
	}

	for i := start; i < len(s); i++ {
		p := s[start : i+1]
		value, _ := mem[p]
		if value == 2 {
			continue
		}
		if value == 1 || isBack(p, mem) {
			path = append(path, p)
			partition_back(s, i+1, path, res, mem)
			path = path[:len(path)-1]
		}
	}
}

func isBack(s string, mem map[string]int) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			mem[s] = 2
			return false
		}
		start++
		end--
	}
	mem[s] = 1
	return true
}

//46
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	permute_back(nums, 0, []int{}, used, &res)
	return res
}

func permute_back(nums []int, index int, path []int, used []bool, res *[][]int) {
	if index == len(nums) {
		newP := make([]int, len(path))
		copy(newP, path)
		*res = append(*res, newP)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			permute_back(nums, index+1, path, used, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

//47
