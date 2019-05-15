// https://leetcode-cn.com/problems/count-complete-tree-nodes/
// 222. 完全二叉树的节点个数
package cpt6

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{1, 2, 3, 4, 5, 6})}, want: 6},
		{name: "test2", args: args{root: CreateTree([]int{1})}, want: 1},
		{name: "test3", args: args{root: CreateTree([]int{1, 2})}, want: 2},
		{name: "test4", args: args{root: CreateTree([]int{1, 2, 3, 4})}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
