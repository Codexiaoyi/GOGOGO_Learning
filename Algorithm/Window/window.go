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

//*******************************5. 最长回文子串*******************
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	mem := make(map[string]struct{})
	res := ""
	start, end := 0, 0
	for start < len(s)-1 {
		end = start
		for end < len(s) {
			cur := string(s[start : end+1])
			if _, ok := mem[cur]; !ok && isPalindrome(cur) && len(cur) > len(res) {
				res = cur
				mem[cur] = struct{}{}
			}
			end++
		}
		start++
	}
	return res
}

func isPalindrome(s string) bool {
	front, back := 0, len(s)-1
	for front <= back {
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}
