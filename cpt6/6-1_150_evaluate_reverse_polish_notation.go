// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
package cpt6

import "strconv"

func evalRPN(tokens []string) int {
	var (
		stack []int
		top   = -1
	)

	for _, s := range tokens {
		if isOpt(s) {
			stack[top-1] = cal(stack[top-1], stack[top], s)
			stack = stack[:top]
			top--
		} else {
			stack = append(stack, sToN(s))
			top++
		}
	}

	return stack[0]
}

func sToN(nstr string) (num int) {
	num, _ = strconv.Atoi(nstr)
	return
}

func isOpt(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/"
}

func cal(num1 int, num2 int, opt string) (result int) {
	switch opt {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return
}
