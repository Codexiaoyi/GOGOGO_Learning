package search

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
