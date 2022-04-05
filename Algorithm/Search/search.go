package search

import (
	"math"
	"sort"
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
