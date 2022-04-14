package StackAndQueue

//387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
