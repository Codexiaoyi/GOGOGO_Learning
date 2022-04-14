package Sort

import "container/heap"

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

//快排 [4,6,2,3,1] [1,6,2,3,4] [1,3,2,4,6]
func sortQuick(nums []int, start int, end int) {
	if start >= end {
		return
	}
	mark := nums[start]
	left, right := start, end //0,4
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

//********************************215. 数组中的第K个最大元素*********************************
func findKthLargest(nums []int, k int) int {
	l := len(nums)
	//先构建大顶堆
	for i := l/2 - 1; i >= 0; i-- {
		sink(nums, i, l)
	}
	count := 0
	for {
		count++
		if count == k {
			return nums[0]
		}
		nums[0], nums[l-count] = nums[l-count], nums[0]
		sink(nums, 0, l-count)
	}
}

//index：需要下沉的元素索引
func sink(nums []int, index int, length int) {
	for {
		left := index*2 + 1
		right := index*2 + 2
		root := index
		if left < length && nums[left] > nums[root] {
			root = left
		}
		if right < length && nums[right] > nums[root] {
			root = right
		}
		if root == index {
			//根节点最大就不动了
			break
		}
		//交换后继续下沉
		nums[root], nums[index] = nums[index], nums[root]
		index = root
	}
}

//********************************347. 前 K 个高频元素*********************************
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//********************************496. 下一个更大元素 I*********************************
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}
	temp := -1
	for i := 0; i < len(nums1); i++ {
		temp = -1
		index := m[nums1[i]]
		for j := index + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				temp = nums2[j]
				break
			}
		}
		nums1[i] = temp
	}
	return nums1
}
