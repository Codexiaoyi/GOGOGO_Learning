package search

import (
	"math"
	"sort"
	"strings"
)

//******二分法******
func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

//******12. 整数转罗马数字******
func intToRoman(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

//******15. 三数之和******
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	l := len(nums)
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}

//******16. 最接近的三数之和******
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	sumAbs := math.Abs(float64(res - target))
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			abs := math.Abs(float64(sum - target))
			if sum == target {
				return sum
			}
			if abs < sumAbs {
				sumAbs = abs
				res = sum
			}
			if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return res
}

//******125. 验证回文串******
func isPalindrome(s string) bool {
	isNumber := func(b byte) bool { return b >= 48 && b <= 57 }
	isTrue := func(b byte) bool { return (b >= 48 && b <= 57) || (b >= 65 && b <= 90) || (b >= 97 && b <= 122) }
	//只区分数字和字母（不区分大小写）
	front, back := 0, len(s)-1
	for front < back {
		if !isTrue(s[front]) {
			front++
			continue
		}
		if !isTrue(s[back]) {
			back--
			continue
		}
		if s[front] == s[back] || (isNumber(s[front]) && isNumber(s[back]) && (s[front]-s[back] == 32 || s[back]-s[front] == 32)) {
			front++
			back--
			continue
		}
		return false
	}
	return true
}

//******151. 颠倒字符串中的单词******
func reverseWords(s string) string {
	res := strings.Builder{}
	start, end := len(s)-1, len(s)-1
	for start >= 0 {
		if s[start] != ' ' && start != len(s)-1 && s[start+1] == ' ' {
			end = start
		}
		if s[start] == ' ' && start != len(s)-1 && s[start+1] != ' ' {
			res.WriteString(string(s[start+1 : end+1]))
			res.WriteString(" ")
		}
		start--
	}
	if s[0] == ' ' {
		resStr := res.String()
		return string(resStr[:res.Len()-1])
	}
	res.WriteString(s[:end+1])
	return res.String()
}

//******31. 下一个排列******
func nextPermutation(nums []int) {
	//先找到最靠右的第一个非降序的位置
	i := len(nums) - 2
	for i >= 0 && nums[i] > nums[i+1] {
		i--
	}
	if i < 0 {
		//全是降序，直接重排返回
		sort.Ints(nums)
		return
	}
	//找到i右边最接近i的数
	j := len(nums) - 1
	for j > i && nums[i] > nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
}
