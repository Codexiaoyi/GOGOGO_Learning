package imooc

//数组问题
//滑动窗口、双指针
//注意边界

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
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}

//27
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//26
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	flag := nums[n-1] + 1
	//重复部分变成flag
	fast, slow := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[slow] {
			nums[i] = flag
		} else {
			slow = i
		}
	}
	slow = 0
	for fast < n {
		if nums[fast] != flag {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//80
func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	count := 1
	flag := nums[n-1] + 1
	//重复部分变成flag
	fast, slow := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[slow] {
			if count >= 2 {
				nums[i] = flag
			}
			count++
		} else {
			slow = i
			count = 1
		}
	}
	slow = 0
	for fast < n {
		if nums[fast] != flag {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//75
func sortColors(nums []int) {
	//[0...zero] [two...len-1]
	zero, two := -1, len(nums)
	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else {
			if nums[i] != 0 {
				return
			}
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}

//88
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	total := m + n - 1         //1
	index1, index2 := m-1, n-1 //0,0
	for {
		if index1 < 0 || index2 < 0 {
			break
		}
		if nums1[index1] > nums2[index2] {
			nums1[total] = nums1[index1]
			index1--
		} else {
			nums1[total] = nums2[index2]
			index2--
		}
		total--
	}
	if index2 >= 0 {
		for i := 0; i <= index2; i++ {
			nums1[i] = nums2[i]
		}
	}
}

//215
func findKthLargest(nums []int, k int) int {
	index := len(nums) - k
	fastSort := func(nums []int, start, end int) int {
		pivot := nums[start]
		for start < end {
			//从右到左找第一个小于pivot的点
			for start < end && nums[end] >= pivot {
				end--
			}
			if start < end {
				//找到小于的点
				nums[start] = nums[end]
				start++
			}
			//从左到右找第一个大于pivot的点
			for start < end && nums[start] <= pivot {
				start++
			}
			if start < end {
				nums[end] = nums[start]
				end--
			}
		}
		nums[start] = pivot
		return start
	}

	resultIndex := fastSort(nums, 0, len(nums)-1)
	for {
		if resultIndex == index {
			return nums[index]
		}
		if resultIndex < index {
			resultIndex = fastSort(nums, resultIndex+1, len(nums)-1)
		}
		if resultIndex > index {
			resultIndex = fastSort(nums, 0, resultIndex-1)
		}
	}
}

//167
func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

//344
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//11
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		minHeight := 0
		isRightMin := height[left] > height[right]
		if isRightMin {
			minHeight = height[right]
		} else {
			minHeight = height[left]
		}
		capacity := (right - left) * minHeight
		if capacity > max {
			max = capacity
		}
		if isRightMin {
			right--
		} else {
			left++
		}
	}
	return max
}

//209
func minSubArrayLen(target int, nums []int) int {
	//滑动窗口，小了就变大
	left, right := 0, -1 //[left...right] 为了一开始没有任何值，所以right取-1
	size, sum := len(nums)+1, 0
	for left < len(nums) {
		if right < len(nums)-1 && sum < target {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= target {
			temp := right - left + 1
			if temp < size {
				size = temp
			}
		}
	}
	if size == len(nums)+1 {
		return 0
	}
	return size
}

//3
//438
//76
