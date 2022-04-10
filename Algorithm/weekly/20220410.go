package weekly

import (
	"container/heap"
	"strconv"
)

//**************1***************
func largestInteger(num int) int {
	str := []byte(strconv.Itoa(num))
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if (str[i]+str[j])%2 == 0 && str[i] < str[j] {
				str[i], str[j] = str[j], str[i]
			}
		}
	}
	res, _ := strconv.Atoi(string(str))
	return res
}

//**************2***************
type IHeap []int

func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	h := IHeap(nums)
	heap.Init(&h)
	for k > 0 {
		x := heap.Pop(&h).(int)
		x++
		heap.Push(&h, x)
		k--
	}
	res := 1
	for _, v := range h {
		res = (res * v) % (1e9 + 7)
	}
	return res
}
