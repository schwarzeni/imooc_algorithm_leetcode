// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
package cpt6

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorNotBST(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{root: CreateTree([]int{3, 5, 1, 6, 2, 0, 8, NULL_NODE, NULL_NODE, 7, 4}), p: &TreeNode{Val: 5}, q: &TreeNode{Val: 1}},
			want: &TreeNode{Val: 3},
		},
		{
			name: "test2",
			args: args{root: CreateTree([]int{3, 5, 1, 6, 2, 0, 8, NULL_NODE, NULL_NODE, 7, 4}), p: &TreeNode{Val: 5}, q: &TreeNode{Val: 4}},
			want: &TreeNode{Val: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestorNotBST(tt.args.root, tt.args.p, tt.args.q); got == nil || !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestorNotBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
