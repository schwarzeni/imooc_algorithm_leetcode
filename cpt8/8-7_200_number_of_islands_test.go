// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿的个数
// 输入:
// 11000
// 11000
// 00100
// 00011
// 输出: 3
package cpt8

import "testing"

var grid1 = [][]byte{
	{'1', '1', '1', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '0', '0', '0'},
}

var grid2 = [][]byte{
	{'1', '1', '0', '0', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
}

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "tes1", args: args{grid: grid1}, want: 1},
		{name: "tes2", args: args{grid: grid2}, want: 3},
		{name: "e1", args: args{grid: [][]byte{{'1', '0', '1', '1', '0', '1', '1'}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
