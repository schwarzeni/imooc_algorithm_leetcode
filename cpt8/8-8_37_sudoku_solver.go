//https://leetcode-cn.com/problems/sudoku-solver/
// 37. 解数独
// 编写一个程序，通过已填充的空格来解决数独问题。
//一个数独的解法需遵循如下规则：
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
//空白格用 '.' 表示。
// 给定的数独序列只包含数字 1-9 和字符 '.' 。
// 你可以假设给定的数独只有唯一解。
// 给定数独永远是 9x9 形式的。
package cpt8

func solveSudoku(board [][]byte) {
	var (
		lmap = [9][10]int{}
		rmap = [9][10]int{}
		gmap = [9][10]int{}

		// ( i / 3 ) + (j / 3) * 3
		// (5, 4) --> 1 + 3 = 4
		// (7, 7) --> 2 + 6 = 9
		getGroupIdx = func(x, y int) int {
			return x/3 + (y/3)*3
		}

		ssolveeSK func(px int, py int) bool
	)
	const (
		TYPE_BLANK  = '.'
		STUT_USED   = 1
		STUT_UNUSED = 0
	)

	for i, y := range board {
		for j, _ := range y {
			var (
				v byte
				c int
			)
			if v = board[i][j]; v != TYPE_BLANK {
				c = STUT_USED
				v = v - '0'
				lmap[i][v] = c
				rmap[j][v] = c
				gmap[getGroupIdx(i, j)][v] = c
			}
		}
	}

	ssolveeSK = func(px int, py int) bool {
		if py == 9 {
			px++
			py = 0
		}
		if px == 9 {
			return true
		}
		if board[px][py] == TYPE_BLANK {
			for i := 1; i < 10; i++ {
				gid := getGroupIdx(px, py)
				if lmap[px][i] != STUT_USED && rmap[py][i] != STUT_USED && gmap[gid][i] != STUT_USED {
					lmap[px][i] = STUT_USED
					rmap[py][i] = STUT_USED
					gmap[gid][i] = STUT_USED
					board[px][py] = byte(i) + '0'
					if ok := ssolveeSK(px, py+1); ok {
						return true
					}
					lmap[px][i] = STUT_UNUSED
					rmap[py][i] = STUT_UNUSED
					gmap[gid][i] = STUT_UNUSED
					board[px][py] = TYPE_BLANK
				}
			}
		} else {
			if ok := ssolveeSK(px, py+1); ok {
				return true
			}
		}
		return false
	}
	ssolveeSK(0, 0)
}
