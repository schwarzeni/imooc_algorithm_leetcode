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

import "testing"

var board = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S', 'G'},
	{'A', 'D', 'E', 'E', 'H'},
}

var board2 = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'E', 'S'},
	{'A', 'D', 'E', 'E'}}

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{board: board, word: "ABCCED"}, want: true},
		{name: "test2", args: args{board: board, word: "SEE"}, want: true},
		{name: "test3", args: args{board: board, word: "ABCB"}, want: false},
		{name: "test4", args: args{board: board, word: "FCCBASADEESE"}, want: true},
		{name: "test5", args: args{board: board, word: "FCCBASF"}, want: false},
		{name: "test6", args: args{board: board, word: "CCSGHEE"}, want: true},
		{name: "e1", args: args{board: board2, word: "ABCESEEEFS"}, want: true},
		{name: "e2", args: args{board: board2, word: "ADEESEF"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
