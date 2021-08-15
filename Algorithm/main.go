package main

import "fmt"

func main() {
	//letterCombinations("23")
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	strs := make([]string, 0)
	for _, digit := range digits {
		switch digit {
		case '2':
			strs = append(strs, "abc")
		case '3':
			strs = append(strs, "def")
		case '4':
			strs = append(strs, "ghi")
		case '5':
			strs = append(strs, "jkl")
		case '6':
			strs = append(strs, "mno")
		case '7':
			strs = append(strs, "pqrs")
		case '8':
			strs = append(strs, "tuv")
		case '9':
			strs = append(strs, "wxyz")
		}
	}
	results := make([]string, 0)
	path := make([]byte, 0)
	letterCombinations_backTracking(0, strs, path, &results)
	return results
}

func letterCombinations_backTracking(k int, strs []string, path []byte, result *[]string) {
	if len(path) == len(strs) {
		newP := make([]byte, len(path))
		copy(newP, path)
		*result = append(*result, string(newP))
		return
	}
	for i := 0; i < len(strs[k]); i++ {
		path = append(path, strs[k][i])
		letterCombinations_backTracking(k+1, strs, path, result)
		path = path[:len(path)-1]
	}
}
