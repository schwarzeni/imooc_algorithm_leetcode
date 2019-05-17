// https://leetcode-cn.com/problems/unique-paths-ii/
// 63. 不同路径 II
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//  网格中的障碍物和空位置分别用 1 和 0 来表示。
// 说明：m 和 n 的值均不超过 100。
// 示例 1:
// 输入:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// 输出: 2
// 解释:
// 3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
package cpt9

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}}, want: 2},
		{name: "test2", args: args{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		}}, want: 0},
		{name: "test3", args: args{[][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}}, want: 0},
		{name: "test4", args: args{[][]int{
			{0, 1, 0},
		}}, want: 0},
		{name: "test5", args: args{[][]int{
			{0},
			{1},
			{0},
		}}, want: 0},
		{name: "test6", args: args{[][]int{
			{1},
		}}, want: 0},
		{name: "test7", args: args{[][]int{
			{0},
		}}, want: 1},
		{name: "test8", args: args{[][]int{
			{0, 0, 0},
		}}, want: 1},
		{name: "test9", args: args{[][]int{
			{0},
			{0},
			{0},
		}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
