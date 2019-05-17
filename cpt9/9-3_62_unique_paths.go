// https://leetcode-cn.com/problems/unique-paths/
// 62. 不同路径
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？
// 说明：m 和 n 的值均不超过 100。
// 示例 1:
// 输入: m = 3, n = 2
// 输出: 3
// 解释:
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向右 -> 向下
// 2. 向右 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向右
// 示例 2:
// 输入: m = 7, n = 3
// 输出: 28
package cpt9

func uniquePaths(m int, n int) int {
	// m w, n h
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	var (
		cache = make([][]int, n)
		lh    = n - 1
		lw    = m - 1
	)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		cache[0][i] = 1
	}
	for i := 0; i < n; i++ {
		cache[i][0] = 1
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			cache[y][x] = cache[y-1][x] + cache[y][x-1]
		}
	}

	return cache[lh][lw]
}
