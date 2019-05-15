// https://leetcode-cn.com/problems/sum-of-left-leaves/
// 计算给定二叉树的所有左叶子之和。
package cpt6

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{3, 9, 20, NULL_NODE, NULL_NODE, 15, 7})}, want: 24},
		{name: "test2", args: args{root: CreateTree([]int{3, NULL_NODE, 20, 15, 7})}, want: 15},
		{name: "test3", args: args{root: CreateTree([]int{3, NULL_NODE, 20, NULL_NODE, 7})}, want: 0},
		{name: "test3", args: args{root: CreateTree([]int{3})}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
