package Sort

import (
	"container/heap"
)

//********************************912. 排序数组*********************************
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

//********************************378. 有序矩阵中第 K 小的元素*********************************
type heap378 []val

type val struct {
	i, j int
	data int
}

func (h heap378) Len() int { return len(h) }
func (h heap378) Less(i, j int) bool {
	return h[i].data < h[j].data
}
func (h heap378) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *heap378) Push(v interface{}) { *h = append(*h, v.(val)) }
func (h *heap378) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &heap378{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, val{i: i, j: 0, data: matrix[i][0]})
	}
	for i := 1; i < k; i++ {
		v := heap.Pop(h).(val)
		if v.j < n-1 {
			heap.Push(h, val{i: v.i, j: v.j + 1, data: matrix[v.i][v.j+1]})
		}
	}
	return heap.Pop(h).(val).data
}

//****************************************659. 分割数组为连续子序列****************************************
func isPossible(nums []int) bool {
	var list [][]int
LOOP:
	for _, n := range nums {
		for i := len(list) - 1; i >= 0; i-- {
			if list[i][len(list[i])-1]+1 == n {
				list[i] = append(list[i], n)
				continue LOOP
			}
		}
		list = append(list, []int{n})
	}

	for _, v := range list {
		if len(v) < 3 {
			return false
		}
	}
	return true
}
