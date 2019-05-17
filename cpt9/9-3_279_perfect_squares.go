// https://leetcode-cn.com/problems/perfect-squares/
// 279. 完全平方数
// 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
// 示例 1:
// 输入: n = 12
// 输出: 3
// 解释: 12 = 4 + 4 + 4.
// 示例 2:
// 输入: n = 13
// 输出: 2
// 解释: 13 = 4 + 9.
package cpt9

// 个人觉得不太适合用DP
func numSquares(n int) int {
	var (
		cache = make([]int, n+1)
		min   = func(arg1, arg2 int) (minNum int) {
			if arg1 == 0 {
				return arg2
			}
			if arg1 > arg2 {
				minNum = arg2
			} else {
				minNum = arg1
			}
			return
		}
	)
	//for i := 1; i <= n; i++ {
	//	cache[i] = i
	//}
	cache[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			cache[i] = min(cache[i], 1+cache[i-j*j])
		}
	}

	return cache[n]
}
