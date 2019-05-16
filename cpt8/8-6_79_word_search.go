// https://leetcode-cn.com/problems/word-search/
// 79. 单词搜索
// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 示例:
// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// 给定 word = "ABCCED", 返回 true.
// 给定 word = "SEE", 返回 true.
// 给定 word = "ABCB", 返回 false.
package cpt8

import "strconv"

var rmap map[string]bool          // 记录该点是否被访问过
var bordermap = make(map[int]int) // 水平方向上的idx的最大值
var xlen int                      // 竖直方向上的idx的最大值

// 速度慢的原因
// 使用map而不是数组做记录，生成key需要时间（我之前以为二维平面可能层次不齐）
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	for px, x := range board { // 统计每个数组的长度
		bordermap[px] = len(x) - 1
	}
	xlen = len(board) - 1
	for px, x := range board {
		for py, y := range x {
			if y != word[0] {
				continue
			}
			rmap = make(map[string]bool)
			visited(px, py)
			if wordsearch(px, py, word, 1, board) {
				return true
			}
			cancelVisit(px, py)
		}
	}
	return false
}

func wordsearch(x int, y int, word string, widx int, board [][]byte) bool {
	if widx >= len(word) {
		return true
	}
	var paths = [][]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}
	for _, p := range paths {
		var (
			x = p[0]
			y = p[1]
		)
		if canVisit(x, y) && board[x][y] == word[widx] && visited(x, y) {
			if wordsearch(x, y, word, widx+1, board) {
				return true
			}
			// TODO: !!!!!! 记得取消访问
			cancelVisit(x, y)
		}
	}
	return false
}

func canVisit(x int, y int) bool {
	return x >= 0 && x <= xlen && y >= 0 && y <= bordermap[x]
}

func visited(x int, y int) bool {
	key := strconv.Itoa(x) + "_" + strconv.Itoa(y)
	if isVisited, ok := rmap[key]; ok && isVisited {
		return false
	}
	rmap[key] = true
	return true
}

func cancelVisit(x int, y int) {
	key := strconv.Itoa(x) + "_" + strconv.Itoa(y)
	rmap[key] = false
}
