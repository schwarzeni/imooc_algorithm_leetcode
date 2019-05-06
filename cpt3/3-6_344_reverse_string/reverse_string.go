// 反转字符串
// https://leetcode-cn.com/problems/reverse-string/
package reverse_string

func reverseString(s []byte) {
	var (
		li = 0
		ri = len(s) - 1
	)

	for li < ri {
		s[li], s[ri] = s[ri], s[li]
		li++
		ri--
	}
}
