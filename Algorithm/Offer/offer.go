package offer

//****************************************剑指 Offer 04. 二维数组中的查找****************************************
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		if matrix[i][col-1] < target {
			continue
		}
		for j := 0; j < col; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//****************************************剑指 Offer 05. 替换空格****************************************
func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		if s == " " {
			return "%20"
		} else {
			return s
		}
	}
	midle := len(s) / 2
	left := replaceSpace(s[:midle])
	right := replaceSpace(s[midle:])
	return left + right
}
