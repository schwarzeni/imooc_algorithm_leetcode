// https://leetcode-cn.com/problems/simplify-path/
// 71. 简化路径
// 请注意，返回的规范路径必须始终以斜杠 / 开头，
// 并且两个目录名之间必须只有一个斜杠 /。
// 最后一个目录名（如果存在）不能以 / 结尾。
// 此外，规范路径必须是表示绝对路径的最短字符串。
package cpt6

import (
	"bytes"
	"strings"
)

func simplifyPath(path string) string {
	var (
		paths  = strings.Split(path, "/")
		stack  []string
		top    = -1
		result bytes.Buffer
	)

	for _, p := range paths {
		switch p {
		case "..":
			if top > -1 {
				stack = stack[:top]
				top--
			}
		case ".":
		case "":
		default:
			stack = append(stack, p)
			top++
		}
	}

	if top == -1 {
		result.WriteString("/")
	} else {
		for _, s := range stack {
			result.WriteString("/")
			result.WriteString(s)
		}
	}

	return result.String()
}
