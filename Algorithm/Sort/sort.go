package Sort

//********************************75. 颜色分类*********************************
func sortColors(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
}

//********************************912. 排序数组*********************************
func sortArray(nums []int) []int {
	return sortMerge(nums)
}

//堆排序
func sortHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		//从最后一个非叶子节点开始
		sortHeap_sink(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 1; i-- {
		//将最后一个节点和堆顶交换
		nums[0], nums[i] = nums[i], nums[0]
		sortHeap_sink(nums, 0, i)
	}
}

func sortHeap_sink(nums []int, index int, length int) {
	for {
		//当前索引下标的左右节点的下标
		left := index*2 + 1
		right := index*2 + 2
		//左右根最大节点的下标
		max_index := index
		if left < length && nums[left] > nums[max_index] {
			max_index = left
		}
		if right < length && nums[right] > nums[max_index] {
			max_index = right
		}
		//根节点大就不下沉
		if max_index == index {
			break
		}
		//根节点小就交换并且继续下沉
		nums[index], nums[max_index] = nums[max_index], nums[index]
		index = max_index
	}
}

//快排
func sortQuick(nums []int, start int, end int) {
	if start >= end {
		return
	}
	mark := nums[start]
	left, right := start, end
	for {
		if left >= right {
			break
		}
		for {
			if left >= right {
				break
			}
			if nums[right] < mark {
				nums[left] = nums[right]
				left++
				break
			}
			right--
		}
		for {
			if left >= right {
				break
			}
			if nums[left] > mark {
				nums[right] = nums[left]
				right--
				break
			}
			left++
		}
	}
	nums[left] = mark
	sortQuick(nums, start, left)
	sortQuick(nums, left+1, end)
}

//归并
func sortMerge(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	midle := len(nums) / 2
	left := sortMerge(nums[:midle])
	right := sortMerge(nums[midle:])
	return sortMerge_merge(left, right)
}

func sortMerge_merge(left, right []int) []int {
	left_index, right_index := 0, 0
	left_length, right_length := len(left), len(right)
	result := make([]int, 0)
	for {
		if left_index >= left_length || right_index >= right_length {
			break
		}
		if left[left_index] <= right[right_index] {
			result = append(result, left[left_index])
			left_index++
		} else {
			result = append(result, right[right_index])
			right_index++
		}
	}

	if left_index != left_length {
		result = append(result, left[left_index:]...)
	}
	if right_index != right_length {
		result = append(result, right[right_index:]...)
	}
	return result
}
