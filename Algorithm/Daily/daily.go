package daily

import (
	"algorithm/List"
	"math"
	"sort"
	"strconv"
	"strings"
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

//*******************************551. 学生出勤记录 I 2021/8/17*******************
func checkRecord(s string) bool {
	late := 0
	absent := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			if absent == 1 {
				return false
			} else {
				absent++
			}
		}
		if s[i] == 'L' {
			late++
			if late == 3 {
				return false
			}
		} else {
			late = 0
		}
	}
	return true
}

//*******************************551. 学生出勤记录 I 2021/8/19*******************
func reverseVowels(s string) string {
	r := make([]rune, 0)
	for _, b := range s {
		if isVowel(b) {
			r = append(r, b)
		}
	}
	s_i := 0
	r_i := len(r) - 1
	s_b := []rune(s)
	for r_i >= 0 {
		if isVowel(s_b[s_i]) {
			s_b[s_i] = rune(r[r_i])
			r_i--
		}
		s_i++
	}
	return string(s_b)
}

func isVowel(b rune) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	return false
}

//*******************************443. 压缩字符串 I 2021/8/21*******************
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	slow, fast := 0, 1
	insert := 0
	length := len(chars)
	for fast <= length {
		//当前字符与前面一个字符不相等
		if fast == length || chars[fast] != chars[fast-1] {
			count := fast - slow
			chars[insert] = chars[fast-1]
			insert++
			//后面跟数字
			if count > 1 {
				count_byte := []byte(strconv.Itoa(count))
				for i := 0; i < len(count_byte); i++ {
					chars[insert] = count_byte[i]
					insert++
				}
			}
			slow = fast
		}
		fast++
	}
	return insert
}

//*******************************797. 所有可能的路径 2021/8/26******************
func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	allPathsSourceTarget_backtrack(graph, 0, path, &result)
	return result
}

func allPathsSourceTarget_backtrack(graph [][]int, index int, path []int, result *[][]int) {
	if len(graph)-1 == index {
		newP := make([]int, len(path)+1)
		copy(newP, path)
		newP[len(newP)-1] = index
		*result = append(*result, newP)
		return
	}

	for i := 0; i < len(graph[index]); i++ {
		path = append(path, index)
		allPathsSourceTarget_backtrack(graph, graph[index][i], path, result)
		path = path[:len(path)-1]
	}
}

//*******************************881. 救生艇 2021/8/27*******************
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	ans := 0
	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			left++
			right--
		}
		ans += 1
	}
	return ans
}

//*******************************1480. 一维数组的动态和 2021/8/28*******************
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

//*******************************1588. 所有奇数长度子数组的和 2021/8/29*******************
func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)
	ans := 0
	width := 1
	for {
		if width > length {
			break
		}
		for i := 0; i+width-1 < length; i++ {
			for j := i; j < width+i; j++ {
				ans += arr[j]
			}
		}
		width += 2
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//*******************************563. 二叉树的坡度 2021/11/18*******************
func findTilt(root *TreeNode) int {
	_, total := findTilt_Back(root, 0)
	return total
}

func findTilt_Back(root *TreeNode, totalTilt int) (int, int) {
	if root == nil {
		return 0, totalTilt
	}
	left, lt := findTilt_Back(root.Left, totalTilt)
	right, rt := findTilt_Back(root.Right, totalTilt)
	totalTilt += (int)(math.Abs(float64(right) - float64(left)))
	return left + right + root.Val, totalTilt + lt + rt
}

//*******************************397. 整数替换 2021/11/19*******************
func integerReplacement(n int) int {
	res := 0
	for n%2 == 0 {
		n = n / 2
		res++
	}
	if n == 1 {
		return res
	}
	add := integerReplacement(n + 1)
	pub := integerReplacement(n - 1)
	if add > pub {
		res += pub + 1
	} else {
		res += add + 1
	}
	return res
}

//*******************************237. 删除链表中的节点 2021/11/2*******************
func deleteNode(node *List.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//*******************************598. 范围求和 II 2021/11/7*******************
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) <= 0 {
		return m * n
	}
	//取交集，取小
	intersect := func(left []int, right []int) []int {
		res := make([]int, 2)
		if left[0] > right[0] {
			res[0] = right[0]
		} else {
			res[0] = left[0]
		}
		if left[1] > right[1] {
			res[1] = right[1]
		} else {
			res[1] = left[1]
		}
		return res
	}
	res := ops[0]
	for i := 1; i < len(ops); i++ {
		res = intersect(res, ops[i])
	}
	return res[0] * res[1]
}

//*******************************661. 图片平滑器 2022/3/24*******************
func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		res[i] = make([]int, len(img[0]))
	}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			sum := img[i][j]
			count := 1
			if i > 0 {
				if j > 0 {
					sum += img[i-1][j-1]
					count++
				}
				sum += img[i-1][j]
				count++
				if j < len(img[i])-1 {
					sum += img[i-1][j+1]
					count++
				}
			}
			if j > 0 {
				sum += img[i][j-1]
				count++
			}
			if j < len(img[i])-1 {
				sum += img[i][j+1]
				count++
			}
			if i < len(img)-1 {
				if j > 0 {
					sum += img[i+1][j-1]
					count++
				}
				sum += img[i+1][j]
				count++
				if j < len(img[i])-1 {
					sum += img[i+1][j+1]
					count++
				}
			}
			res[i][j] = sum / count
		}
	}
	return res
}

//*******************************954. 二倍数对数组 2022/4/1*******************
func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	queue := make([]int, 0)
	for i, a := range arr {
		if len(queue) == 0 {
			queue = append(queue, a)
		} else if queue[0] == 2*a || 2*queue[0] == a {
			queue = queue[1:]
		} else {
			if i > len(arr)/2 {
				return false
			}
			queue = append(queue, a)
		}
	}
	return len(queue) == 0
}

//*******************************796. 旋转字符串 2022/4/7*******************
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = append(m[s[i]], i)
		} else {
			m[s[i]] = []int{i}
		}
	}
	if v, ok := m[goal[0]]; ok {
		for _, index := range v {
			str := string(s[index:]) + string(s[:index])
			if str == goal {
				return true
			}
		}
	}
	return false
}

//*******************************429. N 叉树的层序遍历 2022/4/8*******************

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	l := len(queue)
	for l != 0 {
		floor := make([]int, 0, l)
		for i := 0; i < l; i++ {
			floor = append(floor, queue[i].Val)
			queue = append(queue, queue[i].Children...)
		}
		queue = queue[l:]
		l = len(queue)
		res = append(res, floor)
	}
	return res
}

//*******************************357. 统计各位数字都不同的数字个数 2022/4/11*******************
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + (dp[i-1]-dp[i-2])*(10-(i-1))
	}
	return dp[n]
}

//*******************************1672. 最富有客户的资产总量 2022/4/14*******************
func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		curMax := 0
		for j := 0; j < len(accounts[i]); j++ {
			curMax += accounts[i][j]
		}
		if curMax > max {
			max = curMax
		}
	}
	return max
}

//*******************************819. 最常见的单词 2022/4/17*******************
func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]bool)
	for i := 0; i < len(banned); i++ {
		bannedMap[banned[i]] = true
	}
	isWord := func(b byte) bool {
		return (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
	}
	wordMap := make(map[string]int)
	start, end := 0, 0
	for end < len(paragraph) {
		if isWord(paragraph[end]) && end > 0 && !isWord(paragraph[end-1]) {
			start = end
		}
		if isWord(paragraph[end]) && end < len(paragraph)-1 && !isWord(paragraph[end+1]) {
			word := strings.ToLower(string(paragraph[start : end+1]))
			if !bannedMap[word] {
				wordMap[word] += 1
			}
		}
		end++
	}
	if isWord(paragraph[len(paragraph)-1]) {
		word := strings.ToLower(string(paragraph[start:]))
		if !bannedMap[word] {
			wordMap[word] += 1
		}
	}
	max := 0
	res := ""
	for k, v := range wordMap {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}

//*******************************953. 验证外星语词典 2022/5/18*******************
// func isAlienSorted(words []string, order string) bool {
// 	if len(words) <= 1 {
// 		return true
// 	}
// 	m := make(map[byte]int)
// 	for i := 0; i < len(order); i++ {
// 		m[order[i]] = i
// 	}
// 	length := 0
// 	for i := 0; i < len(words); i++ {
// 		if len(words[i]) > length {
// 			length = len(words[i])
// 		}
// 	}
// 	for i := 0; i < length; i++ {

// 	}
// 	return true
// }
