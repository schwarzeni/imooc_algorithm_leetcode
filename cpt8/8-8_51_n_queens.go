// https://leetcode-cn.com/problems/n-queens/
// 51. N皇后
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 示例:
// 输入: 4
// 输出: [
//  [".Q..",  // 解法 1
//   "...Q",
//   "Q...",
//   "..Q."],
//  ["..Q.",  // 解法 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// 解释: 4 皇后问题存在两个不同的解法。
package cpt8

func solveNQueens(n int) [][]string {
	var (
		col     = make([]int, n)
		dia1    = make([]int, 2*n-1)
		dia2    = make([]int, 2*n-1)
		result  [][]string
		byteArr []byte
		visit   func(x, y int) bool
		unvisit func(x, y int)
		solve   func(level int, strs []string)
	)

	if n < 1 {
		return [][]string{}
	}

	visit = func(x, y int) bool {
		if col[y] == 0 && dia1[x+y] == 0 && dia2[x-y+n-1] == 0 {
			col[y] = 1
			dia1[x+y] = 1
			dia2[x-y+n-1] = 1
			return true
		}
		return false
	}

	unvisit = func(x, y int) {
		col[y] = 0
		dia1[x+y] = 0
		dia2[x-y+n-1] = 0
	}

	solve = func(level int, strs []string) {
		if level == n {
			var arr = make([]string, len(strs))
			copy(arr, strs)
			result = append(result, arr)
			return
		}

		byteArr = make([]byte, n)
		for i := 0; i < n; i++ {
			byteArr[i] = '.'
		}
		for i := 0; i < n; i++ {
			if visit(level, i) {
				byteArr[i] = 'Q'
				solve(level+1, append(strs, string(byteArr)))
				byteArr[i] = '.'
				unvisit(level, i)
			}
		}
	}

	solve(0, []string{})

	return result
}
