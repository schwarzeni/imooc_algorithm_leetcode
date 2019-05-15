// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度
package cpt6

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{3, 9, 20, NULL_NODE, NULL_NODE, 15, 7})}, want: 3},
		{name: "test2", args: args{root: CreateTree([]int{3})}, want: 1},
		{name: "test3", args: args{root: CreateTree([]int{})}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
