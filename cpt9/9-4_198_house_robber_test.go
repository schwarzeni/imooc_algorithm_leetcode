// https://leetcode-cn.com/problems/house-robber/
// 198. 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
// 示例 1:
// 输入: [1,2,3,1]
// 输出: 4
// 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 2:
// 输入: [2,7,9,3,1]
// 输出: 12
// 解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//      偷窃到的最高金额 = 2 + 9 + 1 = 12 。
package cpt9

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{2, 7, 9, 3, 1}}, want: 12},
		{name: "test2", args: args{nums: []int{1, 2, 3, 1}}, want: 4},
		{name: "test3", args: args{nums: []int{7, 2, 9, 3, 1}}, want: 17},
		{name: "test4", args: args{nums: []int{7, 2, 9, 18, 1}}, want: 25},
		{name: "test5", args: args{nums: []int{7, 2, 9, 18, 1}}, want: 25},
		{name: "test6", args: args{nums: []int{7, 2, 9, 18, 100}}, want: 116},
		{name: "test7", args: args{nums: []int{7}}, want: 7},
		{name: "test7", args: args{nums: []int{}}, want: 0},
		{name: "test8", args: args{nums: []int{7, 1}}, want: 7},
		{name: "test9", args: args{nums: []int{1, 7}}, want: 7},
		{name: "test10", args: args{nums: []int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
