// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿的个数
// 输入:
// 11000
// 11000
// 00100
// 00011
// 输出: 3
package cpt8

const (
	TYPE_ISLAND    = '1' // 为岛屿
	TYEP_SEA       = '0' // 为水
	STUS_UNVISITED = 0   // 未访问状态
	STUS_VISITED   = 1   // 访问状态
)

func numIslands(grid [][]byte) int {
	var (
		mmap          [][]byte
		moveDirection = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 顺时针移动
		gridMapInfo   struct {
			y []int
			x int
		}
		initIslansMap   func(grid [][]byte)
		isIsland        func(x int, y int, grid [][]byte) bool
		visitIsland     func(x int, y int)
		isVisitedIsland func(x int, y int) bool
		visitIslands    func(x int, y int, grid [][]byte)
		isInGrid        func(x int, y int) bool
	)

	visitIslands = func(x int, y int, grid [][]byte) {
		for _, d := range moveDirection {
			newX := x + d[0]
			newY := y + d[1]
			if isInGrid(newX, newY) && isIsland(newX, newY, grid) && !isVisitedIsland(newX, newY) {
				visitIsland(newX, newY)
				visitIslands(newX, newY, grid)
			}
		}
	}

	initIslansMap = func(grid [][]byte) {
		for _, y := range grid {
			m := make([]byte, len(y))
			for idx, _ := range m {
				m[idx] = STUS_UNVISITED
			}
			gridMapInfo.y = append(gridMapInfo.y, len(y)-1)
			mmap = append(mmap, m)
		}
		gridMapInfo.x = len(grid) - 1
	}

	isIsland = func(x int, y int, grid [][]byte) bool {
		return grid[x][y] == TYPE_ISLAND
	}

	isVisitedIsland = func(x int, y int) bool {
		return mmap[x][y] == STUS_VISITED
	}

	visitIsland = func(x int, y int) {
		mmap[x][y] = STUS_VISITED
	}

	isInGrid = func(x int, y int) bool {
		return x >= 0 && x <= gridMapInfo.x && y >= 0 && y <= gridMapInfo.y[x]
	}

	var counter = 0
	if len(grid) == 0 {
		return counter
	}
	initIslansMap(grid)
	for px, x := range grid {
		for py, _ := range x {
			if isIsland(px, py, grid) && !isVisitedIsland(px, py) {
				visitIsland(px, py)
				counter++
				visitIslands(px, py, grid)
			}
		}
	}
	return counter

}
