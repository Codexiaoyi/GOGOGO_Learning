package daily

import (
	"sort"
)

//*******************************2021/7/30 171*******************
func titleToNumber(columnTitle string) int {
	res := 0
	multiple := 1
	for i := len(columnTitle); i >= 1; i-- {
		word := columnTitle[i-1] - 'A' + 1
		wordInt := int(word)
		res = res + wordInt*multiple
		multiple *= 26
	}
	return res
}

//*******************************69. x 的平方根*******************
func mySqrt(x int) int {
	count := 1
	for {
		sqrt := count * count
		if sqrt > x {
			return count - 1
		}
		count++
	}
}

//*******************************1337. 矩阵中战斗力最弱的 K 行*******************
func kWeakestRows(mat [][]int, k int) []int {
	type army struct {
		index int
		power int
	}
	armyPower := make([]*army, 0)
	for i := 0; i < len(mat); i++ {
		p := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			p++
		}
		armyPower = append(armyPower, &army{index: i, power: p})
	}
	res := make([]int, k)

	return res
}

//*******************************1337. 矩阵中战斗力最弱的 K 行 2021/8/3*******************
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	//先快排
	quickSort(sortNums, 0, len(sortNums)-1)
	//sort.Ints(sortNums)
	start, end := 0, len(nums)-1
	for {
		//找到最后都没有不同，那就说明原来就是排好序的
		if start == end {
			return 0
		}
		if sortNums[start] != nums[start] {
			break
		}
		start++
	}
	for {
		if start == end {
			return 0
		}
		if sortNums[end] != nums[end] {
			break
		}
		end--
	}
	return end - start + 1
}

func quickSort(arr []int, index_start, index_end int) {
	//定义递归出口
	if index_start >= index_end {
		return
	}
	//快排基准，第一个元素
	num := arr[index_start]
	i, j := index_start, index_end
	for {
		//如果到i,j相等还没找到就算了
		if i >= j {
			break
		}
		//将比基准小的放到基准的左边
		for {
			//i就是需要被换到的位置，j就是当前检测的
			if i >= j {
				break
			}
			if arr[j] < num {
				//当前检测的比基准小就放到i的位置，并把i往前移动，然后去下面的for从前往后
				arr[i] = arr[j]
				i++
				break
			}
			//如果没有找到比基准小的就继续往前找
			j--
		}
		//将比基准大的放到右边
		for {
			if i >= j {
				break
			}
			if arr[i] >= num {
				arr[j] = arr[i]
				j--
				break
			}
			i++
		}
	}
	//全部找完之后，把基准填到坑里
	arr[i] = num
	//左边继续排序
	quickSort(arr, index_start, i)
	//右边继续排序
	quickSort(arr, i+1, index_end)
}

//*******************************611. 有效三角形的个数 2021/8/4*******************
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			result += sort.SearchInts(nums[j+1:], nums[i]+nums[j])
		}
	}
	return result
}

//*******************************802. 找到最终的安全状态 2021/8/5*******************
func eventualSafeNodes(graph [][]int) []int {
	//表示下标i的颜色
	//0表示白色，没有被访问过的
	//1表示灰色，被访问过
	//2表示黑色，已经是安全的
	color := make([]int, len(graph))

	for i, v := range color {
		if v == 0 {
			eventualSafeNodes_visit(i, color, graph)
		}
	}
	res := make([]int, 0)
	for i, v := range color {
		if v == 2 {
			res = append(res, i)
		}
	}
	return res
}

func eventualSafeNodes_visit(currentNode int, color []int, graph [][]int) {
	//当前节点就是下标，与color中下标对应
	//被访问到了，灰色
	color[currentNode] = 1
	if len(graph[currentNode]) == 0 {
		//到终点，黑色
		color[currentNode] = 2
		return
	}
	isSafe := true
	for _, node := range graph[currentNode] {
		//广度遍历
		if color[node] == 0 {
			//没被访问过，就访问
			eventualSafeNodes_visit(node, color, graph)
		}
		if color[node] == 1 {
			//如果都遍历完了，还是灰色，就是不安全
			isSafe = false
		}
	}

	if isSafe {
		//如果当前节点所有的节点遍历完都是安全的，那自己也安全
		color[currentNode] = 2
	}
}

//*******************************1137. 第 N 个泰波那契数 2021/8/8*******************
func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	dp := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		cur := dp[0] + dp[1] + dp[2]
		dp[0], dp[1], dp[2] = dp[1], dp[2], cur
	}
	return dp[2]
}

//*******************************413. 等差数列划分 2021/8/10*******************
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	result := 0
	for i := 3; i <= len(nums); i++ {
		//i是窗口大小
		//起始位置0，终点位置加上窗口大小
		start := 0
		end := start + i - 1
		for end < len(nums) {
			if end == len(nums)-1 {
				if numberOfArithmeticSlices_isArithmetic(nums[start:]) {
					result++
				}
			} else {
				if numberOfArithmeticSlices_isArithmetic(nums[start : end+1]) {
					result++
				}
			}
			start++
			end++
		}
	}
	return result
}

func numberOfArithmeticSlices_isArithmetic(nums []int) bool {
	//nums肯定大于3
	dif := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != dif {
			return false
		}
	}
	return true
}
