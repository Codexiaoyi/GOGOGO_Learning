package days

//****1
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		if res[0] != -1 && res[1] != -1 {
			break
		}
		if nums[left] == target {
			res[0] = left
		} else {
			left++
		}
		if nums[right] == target {
			res[1] = right
		} else {
			right--
		}
	}
	return res
}

//********2  非二分
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		left++
		right--
	}
	return -1
}

//********3
func searchMatrix(matrix [][]int, target int) bool {
	//先二分查找到对应的行
	rleft, rright := 0, len(matrix)-1
	for rleft <= rright {
		rowMid := rleft + (rright-rleft)/2
		if matrix[rowMid][0] <= target && matrix[rowMid][len(matrix[rowMid])-1] >= target {
			row := matrix[rowMid]
			//在这一行二分查找
			left, right := 0, len(row)-1
			for left <= right {
				mid := left + (right-left)/2
				if row[mid] == target {
					return true
				}
				if row[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			return false
		}
		if matrix[rowMid][0] > target {
			//在左边
			rright = rowMid - 1
		} else {
			//在右边
			rleft = rowMid + 1
		}
	}
	return false
}
