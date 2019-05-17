// https://leetcode-cn.com/problems/unique-paths-ii/
// 63. 不同路径 II
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//  网格中的障碍物和空位置分别用 1 和 0 来表示。
// 说明：m 和 n 的值均不超过 100。
// 示例 1:
// 输入:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// 输出: 2
// 解释:
// 3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
package cpt9

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// m w, n h
	var (
		cache [][]int
		lh    int
		lw    int
		m     int
		n     = len(obstacleGrid)
	)
	if n == 0 {
		return 0
	}
	if m = len(obstacleGrid[0]); m == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}

	cache = make([][]int, n)
	lh = n - 1
	lw = m - 1

	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
	}
	cache[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[0][i] != 1 {
			cache[0][i] = 1
		} else {
			break
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[i][0] != 1 {
			cache[i][0] = 1
		} else {
			break
		}
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if obstacleGrid[y][x] != 1 {
				cache[y][x] = cache[y-1][x] + cache[y][x-1]
			}
		}
	}

	return cache[lh][lw]
}
