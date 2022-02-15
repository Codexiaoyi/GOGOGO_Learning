package imooc

import (
	"sort"
	"strconv"
)

//查找问题
//查找有无：常用集合
//查找对应关系：常用map字典

//242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] += 1
		m[t[i]] -= 1
	}

	for _, mItem := range m {
		if mItem != 0 {
			return false
		}
	}

	return true
}

//202
func isHappy(n int) bool {
	repeat := make(map[int]struct{})
	for {
		ns := strconv.Itoa(n)
		intArray := make([]int, 0)
		for _, nb := range ns {
			i, _ := strconv.Atoi(string(nb))
			intArray = append(intArray, i)
		}
		sum := 0
		for _, i := range intArray {
			sum += i * i
		}
		n = sum
		if n == 1 {
			return true
		}
		if _, ok := repeat[n]; ok {
			return false
		}
		repeat[n] = struct{}{}
	}
}

//290
//205
//451
func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		m[r] += 1
	}
	is := make([]int, len(m))
	mr := make(map[int]string)
	for r, i := range m {
		if _, ok := mr[i]; !ok {
			is = append(is, i)
		}
		for j := 0; j < i; j++ {
			mr[i] += string(r)
		}
	}
	sort.Ints(is)
	res := ""
	for _, i := range is {
		res = mr[i] + res
	}
	return res
}

//1
//15
//18
//16
