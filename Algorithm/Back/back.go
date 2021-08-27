package back

//*******************************77.组合******************
func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	combine_backTracking(n, k, 1, path, &result)
	return result
}

func combine_backTracking(n, k, start int, path []int, result *[][]int) {
	if len(path) == k {
		newP := make([]int, k)
		copy(newP, path)
		*result = append(*result, newP)
		return
	}
	for i := start; i < n+1; i++ {
		path = append(path, i)
		combine_backTracking(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}

//*******************************17. 电话号码的字母组合******************
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
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

//*******************************22. 括号生成******************
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	path := []byte{}
	generateParenthesis_backTracking(n, 0, path, &result)
	return result
}

func generateParenthesis_backTracking(n, k int, path []byte, result *[]string) {
	if n*2 == k {
		if generateParenthesis_isValid(path) {
			*result = append(*result, string(path))
		}
		return
	}
	path = append(path, '(')
	generateParenthesis_backTracking(n, k+1, path, result)
	path = path[:len(path)-1]
	path = append(path, ')')
	generateParenthesis_backTracking(n, k+1, path, result)
	path = path[:len(path)-1]
}

func generateParenthesis_isValid(s []byte) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

//*******************************79. 单词搜索******************
func exist(board [][]byte, word string) bool {
	visitor := make([][]int, len(board))
	starts := make([][2]int, 0)
	for i := 0; i < len(visitor); i++ {
		visitor[i] = make([]int, len(board[0]))
		for j := 0; j < len(visitor[i]); j++ {
			visitor[i][j] = 0
			if board[i][j] == word[0] {
				//找到起始点
				starts = append(starts, [2]int{i, j})
			}
		}
	}
	for i := 0; i < len(starts); i++ {
		visitor[starts[i][0]][starts[i][1]] = 1
		if exist_backTracking(board, visitor, starts[i], word[1:]) {
			//某个起始点开始的最后找到了
			return true
		}
		visitor[starts[i][0]][starts[i][1]] = 0
	}
	//全都没找到
	return false
}

func exist_backTracking(board [][]byte, visitor [][]int, start [2]int, word string) bool {
	start_row, start_col := start[0], start[1]
	//找到路径
	if len(word) == 0 {
		return true
	}
	//找到周围所有的可以继续走的格子
	starts := make([][2]int, 0)
	//左
	if start_row-1 >= 0 && visitor[start_row-1][start_col] == 0 && board[start_row-1][start_col] == word[0] {
		starts = append(starts, [2]int{start_row - 1, start_col})
	}
	//右
	if start_row+1 < len(board) && visitor[start_row+1][start_col] == 0 && board[start_row+1][start_col] == word[0] {
		starts = append(starts, [2]int{start_row + 1, start_col})
	}
	//上
	if start_col-1 >= 0 && visitor[start_row][start_col-1] == 0 && board[start_row][start_col-1] == word[0] {
		starts = append(starts, [2]int{start_row, start_col - 1})
	}
	//下
	if start_col+1 < len(board[0]) && visitor[start_row][start_col+1] == 0 && board[start_row][start_col+1] == word[0] {
		starts = append(starts, [2]int{start_row, start_col + 1})
	}
	for i := 0; i < len(starts); i++ {
		visitor[starts[i][0]][starts[i][1]] = 1
		if exist_backTracking(board, visitor, starts[i], word[1:]) {
			return true
		}
		visitor[starts[i][0]][starts[i][1]] = 0
	}
	return false
}

//*******************************797. 所有可能的路径******************
func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	allPathsSourceTarget_backtrack(graph, 0, path, &result)
	return result
}

func allPathsSourceTarget_backtrack(graph [][]int, index int, path []int, result *[][]int) {
	if len(graph)-1 == index {
		newP := make([]int, len(path)+1)
		copy(newP, path)
		newP[len(newP)-1] = index
		*result = append(*result, newP)
		return
	}

	for i := 0; i < len(graph[index]); i++ {
		path = append(path, index)
		allPathsSourceTarget_backtrack(graph, graph[index][i], path, result)
		path = path[:len(path)-1]
	}
}
