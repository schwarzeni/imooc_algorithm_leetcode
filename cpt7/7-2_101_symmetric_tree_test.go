// https://leetcode-cn.com/problems/symmetric-tree/
// 101. 对称二叉树
package cpt6

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{root: CreateTree([]int{1, 2, 2, 3, 4, 4, 3})}, want: true},
		{name: "test2", args: args{root: CreateTree([]int{1, 2, 2, NULL_NODE, 3, NULL_NODE, 3})}, want: false},
		{name: "test3", args: args{root: CreateTree([]int{1, 2, 2, NULL_NODE, NULL_NODE, NULL_NODE, 3})}, want: false},
		{name: "test4", args: args{root: CreateTree([]int{1, 2, 2, 3, 4, 3, 3})}, want: false},
		{name: "test5", args: args{root: CreateTree([]int{})}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
