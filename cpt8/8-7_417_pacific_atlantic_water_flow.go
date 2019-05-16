// https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
// 417. 太平洋大西洋水流问题
//
// 给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
//规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
//请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
//提示：
//输出坐标的顺序不重要
//m 和 n 都小于150
//示例：
//给定下面的 5x5 矩阵:
//
//  太平洋 ~   ~   ~   ~   ~
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *
//          *   *   *   *   * 大西洋
//返回:
//[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).
package cpt8

import "sync"

//const (
//	// TODO 将逻辑抽出
//	STUT_UNVISITED  = 4
//	STUT_VISITED    = 3
//	STUT_EXCLUDE    = 2
//	STUT_SELECTED   = 1
//	STUT_UNSELECTED = 0
//)
//
//func pacificAtlantic(matrix [][]int) [][]int {
//	var (
//		result      [][]int
//		selectedMap [][]int
//		mxl         = len(matrix) // 矩阵的x轴长度
//		myl         int           // 矩阵的y轴长度
//
//		isPacific  func(x, y int) bool
//		isAtlantic func(x, y int) bool
//		isSelected func(x, y int) bool
//		//isExcluded func(x, y int) bool
//		//isVisited      func(x, y int) bool
//		selectPoint func(x, y int)
//		//excludePoint   func(x, y int)
//		canAccess      func(x, y, curr int) bool
//		traverseMatrix func(x, y int) (canAccessPacific bool, canAccessAtlantic bool)
//		moveDirection  = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
//	)
//
//	if mxl <= 0 {
//		return [][]int{} // TODO:空矩阵的形式
//	}
//	if myl = len(matrix[0]); myl <= 0 {
//		return [][]int{}
//	}
//	for i := 0; i < mxl; i++ {
//		var a []int
//		for j := 0; j < myl; j++ {
//			a = append(a, STUT_UNSELECTED)
//		}
//		selectedMap = append(selectedMap, a)
//	}
//
//	isPacific = func(x, y int) bool {
//		return y < 0 || x < 0
//	}
//
//	isAtlantic = func(x, y int) bool {
//		return y >= myl || x >= mxl
//	}
//
//	isSelected = func(x, y int) bool {
//		return selectedMap[x][y] == STUT_SELECTED
//	}
//
//	//isExcluded = func(x, y int) bool {
//	//	return selectedMap[x][y] == STUT_EXCLUDE
//	//}
//
//	//isVisited = func(x, y int) bool {
//	//	return selectedMap[x][y] == STUT_VISITED
//	//}
//
//	selectPoint = func(x, y int) {
//		result = append(result, []int{x, y})
//		selectedMap[x][y] = STUT_SELECTED
//	}
//
//	//excludePoint = func(x, y int) {
//	//	selectedMap[x][y] = STUT_EXCLUDE
//	//}
//
//	canAccess = func(x, y, curr int) bool {
//		return matrix[x][y] <= curr && matrix[x][y] >= 0
//	}
//
//	traverseMatrix = func(x, y int) (canAccessPacific bool, canAccessAtlantic bool) {
//		var (
//			newX int
//			newY int
//		)
//		if isSelected(x, y) {
//			return true, true
//		}
//		//if isExcluded(x, y) {
//		//	return false, false
//		//}
//		for _, d := range moveDirection {
//			newX = x + d[0]
//			newY = y + d[1]
//			if isPacific(newX, newY) {
//				canAccessPacific = true
//				continue
//			}
//			if isAtlantic(newX, newY) {
//				canAccessAtlantic = true
//				continue
//			}
//			//if !canAccess(newX, newY, matrix[x][y]) || isExcluded(newX, newY) || isVisited(newX, newY) {
//			//if !canAccess(newX, newY, matrix[x][y]) || isExcluded(newX, newY) {
//			if !canAccess(newX, newY, matrix[x][y]) {
//				continue
//			}
//			if isSelected(newX, newY) {
//				canAccessAtlantic = true
//				canAccessPacific = true
//				break
//			}
//			//prevStut := selectedMap[x][y]
//			//selectedMap[x][y] = STUT_VISITED
//			//cp, ca := traverseMatrix(newX, newY)
//			//selectedMap[x][y] = prevStut
//
//			prev := matrix[x][y]
//			matrix[x][y] = -1
//			cp, ca := traverseMatrix(newX, newY)
//			matrix[x][y] = prev
//
//			if cp {
//				canAccessPacific = true
//			}
//			if ca {
//				canAccessAtlantic = true
//			}
//			if canAccessPacific && canAccessAtlantic {
//				break
//			}
//		}
//		if canAccessAtlantic && canAccessPacific {
//			selectPoint(x, y)
//		}
//		//if !canAccessPacific && !canAccessAtlantic {
//		//	excludePoint(x, y)
//		//}
//		return
//	}
//
//	for i, y := range matrix {
//		for j, _ := range y {
//			traverseMatrix(i, j)
//		}
//	}
//
//	return result
//}

// 另一个思路：从太平洋开始找能到大西洋的点，再从大西洋开始找能到太平洋的点，在取一个交集

func DFS(matrix [][]int, visited *[150][150]bool, pre, i, j int) {
	m, n := len(matrix), len(matrix[0])
	if i < 0 || i >= m || j < 0 || j >= n || pre > matrix[i][j] || visited[i][j] {
		return
	}
	visited[i][j] = true
	DFS(matrix, visited, matrix[i][j], i-1, j)
	DFS(matrix, visited, matrix[i][j], i, j-1)
	DFS(matrix, visited, matrix[i][j], i+1, j)
	DFS(matrix, visited, matrix[i][j], i, j+1)
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	res := make([][]int, 0)
	pacific := [150][150]bool{}
	atlantic := [150][150]bool{}
	m, n := len(matrix), len(matrix[0])

	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i := 0; i < m; i++ {
			DFS(matrix, &pacific, matrix[i][0], i, 0)
			DFS(matrix, &atlantic, matrix[i][n-1], i, n-1)
		}
		wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			DFS(matrix, &pacific, matrix[0][i], 0, i)
			DFS(matrix, &atlantic, matrix[m-1][i], m-1, i)
		}
		wg.Done()
	}(&wg)
	wg.Wait()

	for i := 0; i < 150; i++ {
		for j := 0; j < 150; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
