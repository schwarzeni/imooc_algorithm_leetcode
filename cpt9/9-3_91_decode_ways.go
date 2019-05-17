// https://leetcode-cn.com/problems/decode-ways/
// 91. 解码方法
// 一条包含字母 A-Z 的消息通过以下方式进行了编码：
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
// 示例 1:
// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2:
// 输入: "226"
// 输出: 3
// 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
package cpt9

// v1 s表示的数作为int会产生溢出的情况
//func numDecodings(s string) int {
//	var (
//		snum  int
//		cache []int
//	)
//
//	for _, c := range s {
//		if c-'0' == 0 {
//			return 0
//		}
//	}
//
//	if snum, _ = strconv.Atoi(s); snum == 0 {
//		return 0
//	}
//	cache = make([]int, snum+1)
//	for i := 1; i <= 26 && i <= snum; i++ {
//		cache[i] = 1
//	}
//
//	for i := 2; i <= snum; i++ {
//		cNum := i
//		step := 1
//		for {
//			if cNum >= 1 && cNum <= 26 {
//				cache[i] = cache[i-cNum*step] + cache[i]
//			}
//			step *= 10
//			if cNum = i / step; cNum == 0 {
//				break
//			}
//		}
//	}
//	return cache[snum]
//}

// v2 真正的动态规划，类似于斐波那契数列解法
func numDecodings(s string) int {
	var (
		ssize = len(s)
		d1    = 1
		d2    = 1
	)
	if ssize == 0 || s[0] == '0' {
		return 0
	}

	for i := 0; i < ssize; i++ {
		var tmp = 0
		if s[i] != '0' {
			tmp += d2
		}
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')) {
			tmp += d1
		}
		d1 = d2
		d2 = tmp
	}
	return d2
}
