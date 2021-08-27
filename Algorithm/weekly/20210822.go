package weekly

import (
	"sort"
)

//***********1*********
func findGCD(nums []int) int {
	sort.Ints(nums)
	res := 1
	for i := 1; i <= nums[0]; i++ {
		if nums[0]%i == 0 && nums[len(nums)-1]%i == 0 {
			res = i
		}
	}
	return res
}

//***********2*********
func findDifferentBinaryString(nums []string) string {
	numMap := make(map[string]struct{})
	for _, number := range nums {
		numMap[number] = struct{}{}
	}
	s := ""
	res := ""
	back(numMap, len(nums), &s, &res)
	return res
}

func back(m map[string]struct{}, length int, path, res *string) {
	if *res != "" {
		return
	}
	if len(*path) == length {
		if _, ok := m[*path]; !ok {
			*res = *path
			return
		}
		return
	}
	*path += "0"
	back(m, length, path, res)
	*path = (*path)[:len(*path)-1]
	*path += "1"
	back(m, length, path, res)
	*path = (*path)[:len(*path)-1]
}
