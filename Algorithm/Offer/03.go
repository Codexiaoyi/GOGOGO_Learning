package offer

//****************************************数组中重复的数字****************************************

func findRepeatNumber(nums []int) int {
	cmap := make(map[int]int)
	for _, num := range nums {
		if cmap[num] == 1 {
			return num
		}
		cmap[num] = 1
	}

	return -1
}
