// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// 111. 二叉树的最小深度
package cpt6

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{3, 9, 20, NULL_NODE, NULL_NODE, 15, 7})}, want: 2},
		{name: "test2", args: args{root: CreateTree([]int{3, NULL_NODE, 20, NULL_NODE, 15, NULL_NODE, 7})}, want: 4},
		{name: "test3", args: args{root: CreateTree([]int{3})}, want: 1},
		{name: "test4", args: args{root: CreateTree([]int{})}, want: 0},
		{name: "e1", args: args{root: CreateTree([]int{1, 2})}, want: 2},
		{name: "test5", args: args{root: CreateTree([]int{1, NULL_NODE, 2})}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
