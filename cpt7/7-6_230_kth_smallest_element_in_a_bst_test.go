// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 230. 二叉搜索树中第K小的元素
package cpt6

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{3, 1, 4, NULL_NODE, 2}), k: 1}, want: 1},
		{name: "test2", args: args{root: CreateTree([]int{5, 3, 6, 2, 4, NULL_NODE, NULL_NODE, 1}), k: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
