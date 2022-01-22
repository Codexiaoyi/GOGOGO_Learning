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
