package StackAndQueue

import "strconv"

//********************************150. 逆波兰表达式求值*********************************
/*逆波兰表达式：

逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。

平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
逆波兰表达式主要有以下两个优点：

去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

//********************************394. 字符串解码*********************************
func decodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			//往前找
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				//弹出
				stack = stack[:len(stack)-1]
				//放到临时栈中
				temp = append(temp, v)
			}
			//弹出'['
			stack = stack[:len(stack)-1]
			num := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num = append(num, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			countStr := make([]byte, 0)
			for len(num) != 0 {
				countStr = append(countStr, num[len(num)-1])
				num = num[:len(num)-1]
			}
			count, _ := strconv.Atoi(string(countStr))
			for i := 0; i < count; i++ {
				//按次数弹出临时栈，放入stack中
				for j := 1; j <= len(temp); j++ {
					stack = append(stack, temp[len(temp)-j])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

//********************************133. 克隆图*********************************
