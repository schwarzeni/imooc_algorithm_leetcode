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

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		//{name: "test1", args: args{board: [][]byte{
		//	{'X', 'X', 'X', 'X'},
		//	{'X', 'O', 'O', 'X'},
		//	{'X', 'X', 'O', 'X'},
		//	{'X', 'O', 'X', 'X'},
		//}}},
		//{name: "test2", args: args{board: [][]byte{
		//	{'O', 'O', 'O', 'O'},
		//	{'O', 'O', 'O', 'O'},
		//	{'O', 'O', 'O', 'O'},
		//	{'O', 'O', 'O', 'O'},
		//}}},
		{name: "e1", args: args{board: [][]byte{
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
			{'O', 'O', 'O'},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			fmt.Println(tt.args.board)
		})
	}
}
