package imooc

//二分查找法
func binary_search(input []int, target int) int {
	length := len(input)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if input[mid] == target {
			return mid
		}
		if input[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//283
func move_zeros() {

}

//27

//26

//80
