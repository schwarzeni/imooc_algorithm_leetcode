// https://leetcode-cn.com/problems/balanced-binary-tree/
// 110. 平衡二叉树
package cpt6

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{root: CreateTree([]int{3, 9, 20, NULL_NODE, NULL_NODE, 15, 7})}, want: true},
		{name: "test2", args: args{root: CreateTree([]int{1, 2, 2, 3, 3, NULL_NODE, NULL_NODE, 4, 4})}, want: false},
		{name: "test3", args: args{root: CreateTree([]int{1, 2, 2})}, want: true},
		{name: "test4", args: args{root: CreateTree([]int{1})}, want: true},
		{name: "test5", args: args{root: CreateTree([]int{})}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
