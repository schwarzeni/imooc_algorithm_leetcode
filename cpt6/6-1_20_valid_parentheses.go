// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号
// 输入: "([)]"
// 输出: false
package cpt6

func isValid(s string) bool {
	var (
		size  = len(s)
		lidx  = 0
		stack []uint8
		top   = -1
	)

	if size == 0 {
		return true
	}

	if size%2 != 0 {
		return false
	}

	for lidx < size {
		if isLeft(s[lidx]) {
			stack = append(stack, s[lidx])
			top++
		} else {
			if top == -1 {
				return false
			}
			c := stack[top]
			if !isPair(c, s[lidx]) {
				return false
			} else {
				stack = stack[:top]
				top--
			}
		}
		lidx++
	}

	if top == -1 {
		return true
	} else {
		return false
	}
}

func isLeft(l uint8) bool {
	return l == '(' || l == '{' || l == '['
}

// 考虑使用map
func isPair(l uint8, r uint8) bool {
	return (l == '(' && r == ')') ||
		(l == '[' && r == ']') ||
		(l == '{' && r == '}')
}
