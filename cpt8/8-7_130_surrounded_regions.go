// https://leetcode-cn.com/problems/surrounded-regions/
// 130. 被围绕的区域
// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 示例:
// X X X X
// X O O X
// X X O X
// X O X X
// 运行你的函数后，矩阵变为：
// X X X X
// X X X X
// X X X X
// X O X X
package cpt8

func solve(board [][]byte) {
	const (
		STUS_VISITED     = 1
		STUS_UNVISITED   = 0
		TYPE_TARGET      = 'O'
		TYPE_CHANGE_TO   = 'X'
		TYPE_SAFE_TARGET = '='
	)
	var (
		borderTargetMap [][]int
		innerTargetMap  [][]int
		isVisitedMap    [][]int
		boardX          = len(board)
		boardY          int
		isTarget        func(x, y int) bool
		isValid         func(x, y int) bool
		isBorder        func(x, y int) bool
		isVisited       func(x, y int) bool
		isSafe          func(x, y int) bool
		visit           func(x, y int)
		moveDirection   = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	)

	isTarget = func(x, y int) bool {
		return board[x][y] == TYPE_TARGET
	}

	isValid = func(x, y int) bool {
		return x >= 0 && x < boardX && y >= 0 && y < boardY
	}

	isBorder = func(x, y int) bool {
		return x == 0 || y == 0 || x == boardX-1 || y == boardY-1
	}

	isVisited = func(x, y int) bool {
		return isVisitedMap[x][y] == STUS_VISITED
	}

	visit = func(x, y int) {
		isVisitedMap[x][y] = STUS_VISITED
	}

	isSafe = func(x, y int) bool {
		return board[x][y] == TYPE_SAFE_TARGET
	}

	// 进行一些初始化操作
	if boardX < 3 {
		return
	}
	boardY = len(board[0])
	if boardY < 3 {
		return
	}
	isVisitedMap = make([][]int, boardX)
	for idx, _ := range isVisitedMap {
		var y []int
		for i := 0; i < boardY; i++ {
			y = append(y, STUS_UNVISITED)
		}
		isVisitedMap[idx] = y
	}

	// 找到所有在边缘的'O'
	for y := 0; y < boardY; y++ {
		if isTarget(0, y) {
			borderTargetMap = append(borderTargetMap, []int{0, y})
		}
		if isTarget(boardX-1, y) {
			borderTargetMap = append(borderTargetMap, []int{boardX - 1, y})
		}
	}
	for x := 1; x < boardX-1; x++ {
		if isTarget(x, 0) {
			borderTargetMap = append(borderTargetMap, []int{x, 0})
		}
		if isTarget(x, boardY-1) {
			borderTargetMap = append(borderTargetMap, []int{x, boardY - 1})
		}
	}

	for {
		var (
			qsize int
			px    int
			py    int
			currP []int
		)
		if qsize = len(borderTargetMap); qsize == 0 {
			break
		}
		currP = borderTargetMap[0]
		borderTargetMap = borderTargetMap[1:]

		// TODO: e1 错误点：为考虑边界点相邻的情况
		//if !isVisited(currP[0], currP[1]) {
		//visit(currP[0], currP[1])
		//} else {
		//	continue
		//}
		visit(currP[0], currP[1])

		for _, d := range moveDirection {
			px = currP[0] + d[0]
			py = currP[1] + d[1]
			if isValid(px, py) && isTarget(px, py) && !isVisited(px, py) {
				if !isBorder(px, py) {
					innerTargetMap = append(innerTargetMap, []int{px, py})
					borderTargetMap = append(borderTargetMap, []int{px, py})
				} else {
					visit(px, py)
				}
			}
		}
	}

	for _, v := range innerTargetMap {
		board[v[0]][v[1]] = TYPE_SAFE_TARGET
	}

	for i := 1; i < boardX-1; i++ {
		for j := 1; j < boardY-1; j++ {
			if isSafe(i, j) {
				board[i][j] = TYPE_TARGET
			} else if isTarget(i, j) {
				board[i][j] = TYPE_CHANGE_TO
			}
		}
	}
}
