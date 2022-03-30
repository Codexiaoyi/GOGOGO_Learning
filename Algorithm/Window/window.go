package window

//*******************************3. 无重复字符的最长子串*******************
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	start, end := 0, 0
	maxLength := 0
	m := make(map[byte]bool)
	for start < len(s) {
		if start != 0 {
			delete(m, s[start-1])
		}
		for end < len(s) && !m[s[end]] {
			m[s[end]] = true
			end++
		}
		cur := end - start
		if cur > maxLength {
			maxLength = cur
		}
		start++
	}
	return maxLength
}
