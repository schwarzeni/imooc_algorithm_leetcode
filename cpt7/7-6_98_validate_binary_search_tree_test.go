// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树
package cpt6

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{root: CreateTree([]int{2, 1, 3})}, want: true},
		{name: "test2", args: args{root: CreateTree([]int{5, 1, 4, NULL_NODE, 3, 6})}, want: false},
		{name: "e1", args: args{root: CreateTree([]int{10, 5, 15, NULL_NODE, NULL_NODE, 6, 20})}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
