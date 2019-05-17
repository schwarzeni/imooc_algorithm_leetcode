// https://leetcode-cn.com/problems/minimum-path-sum/
// 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例:
// 输入:
// [
// [1,3,1],
// [1,5,1],
// [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。
package cpt9

import "math"

func minPathSum(grid [][]int) int {
	var (
		gw    int         // 总宽度
		gh    = len(grid) // 总高度
		walk  func(level int, idx int) int
		cache [][]int
	)
	if gh == 0 {
		return 0
	}
	if gw = len(grid[0]); gw == 0 {
		return 0
	}

	for i := 0; i < gh; i++ {
		cache = append(cache, make([]int, gw))
	}
	for i := 0; i < gh; i++ {
		for j := 0; j < gw; j++ {
			cache[i][j] = math.MaxInt64
		}
	}

	walk = func(level int, idx int) int {
		if level == gh-1 && idx == gw-1 { // 终止条件
			return grid[level][idx]
		}
		var (
			rl = math.MaxInt64 // 向右最小值
			bl = math.MaxInt64 // 向下最小值
			ml int             // 比较后的最小值
		)
		if idx+1 < gw { // 向右
			if rl = cache[level][idx+1]; rl == math.MaxInt64 {
				rl = walk(level, idx+1)
				cache[level][idx+1] = rl
			}
		}
		if level+1 < gh {
			if bl = cache[level+1][idx]; bl == math.MaxInt64 {
				bl = walk(level+1, idx)
				cache[level+1][idx] = bl
			}
		}
		if bl < rl {
			ml = bl
		} else {
			ml = rl
		}
		return ml + grid[level][idx]
	}

	return walk(0, 0)

}

// v2 动态规划
//func minPathSum(grid [][]int) int {
//	m := len(grid)
//	if m == 0 {
//		return 0
//	}
//	n := len(grid[0])
//	if n == 0 {
//		return 0
//	}
//	for i := 1; i < n; i++ {
//		grid[0][i] += grid[0][i-1]
//	}
//	for i := 1; i < m; i++ {
//		grid[i][0] += grid[i-1][0]
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
//		}
//	}
//	return grid[m-1][n-1]
//}
//
//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
