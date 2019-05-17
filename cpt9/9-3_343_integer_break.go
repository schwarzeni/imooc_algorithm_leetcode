// https://leetcode-cn.com/problems/integer-break/
// 343. 整数拆分
// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
// 示例 1:
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1。
// 示例 2:
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
// 说明: 你可以假设 n 不小于 2 且不大于 58。
package cpt9

// v1 自顶向下的解法
//func integerBreak(n int) int {
//	var (
//		walk  func(n int) int
//		cache = make([]int, n)
//	)
//
//	if n == 2 {
//		return 1
//	}
//
//	walk = func(step int) int {
//		var (
//			max = 0
//		)
//		if step == 1 {
//			return 1
//		}
//		for i := step - 1; i >= 1; i-- {
//			var (
//				v      int
//				remain = step - i
//			)
//
//			if v = cache[remain]; v == 0 {
//				v = walk(remain)
//				if v < remain { // 如果拆了比不拆还小，则用原值覆盖之前的结果
//					v = remain
//				}
//				cache[remain] = v
//			}
//			if v*i > max {
//				max = v * i
//			}
//		}
//		return max
//	}
//
//	return walk(n)
//}

// 自底向上的解法 动态规划
func integerBreak(n int) int {
	var (
		cache = make([]int, n+1)
		max3  = func(i int, i2 int, i3 int) int {
			if i > i2 && i > i3 {
				return i
			} else if i2 > i && i2 > i3 {
				return i2
			} else {
				return i3
			}
		}
	)

	cache[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			cache[i] = max3(cache[i], j*(i-j), j*cache[i-j])
		}
	}
	return cache[n]
}
