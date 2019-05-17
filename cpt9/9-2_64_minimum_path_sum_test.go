// https://leetcode-cn.com/problems/minimum-path-sum/
// 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例:
// 输入:
// [
// [1,3,1],
// [1,5,1],
// [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。
package cpt9

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "tes1", args: args{grid: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}}, want: 7},
		{name: "tes2", args: args{grid: [][]int{
			{1},
		}}, want: 1},
		{name: "tes3", args: args{grid: [][]int{
			{},
		}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
