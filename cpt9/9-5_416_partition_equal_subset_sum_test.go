// https://leetcode-cn.com/problems/partition-equal-subset-sum/
// 416. 分割等和子集
// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//注意:
//每个数组中的元素不会超过 100
//数组的大小不会超过 200
//示例 1:
//输入: [1, 5, 11, 5]
//输出: true
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//示例 2:
//输入: [1, 2, 3, 5]
//输出: false
//解释: 数组不能分割成两个元素和相等的子集.
package cpt9

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{[]int{1, 5, 11, 5}}, want: true},
		{name: "test2", args: args{[]int{1, 2, 3, 5}}, want: false},
		{name: "err1", args: args{[]int{1, 3, 5, 5, 5, 5}}, want: false},
		{name: "err2", args: args{[]int{1, 5, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
