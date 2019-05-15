// https://leetcode-cn.com/problems/path-sum/
// 112. 路径总和
// 说明: 叶子节点是指没有子节点的节点。
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \      \
//         7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
package cpt6

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{root: CreateTree([]int{5, 4, 8, 11, NULL_NODE, 13, 4, 7, 2, NULL_NODE, NULL_NODE, NULL_NODE, 4}), sum: 22}, want: true},
		{name: "test2", args: args{root: CreateTree([]int{5, 4, 8, 11, NULL_NODE, 13, 4, 7, 2, NULL_NODE, NULL_NODE, NULL_NODE, 4}), sum: 24}, want: false},
		{name: "test3", args: args{root: CreateTree([]int{5}), sum: 3}, want: false},
		{name: "test3", args: args{root: CreateTree([]int{5}), sum: 5}, want: true},
		{name: "test3", args: args{root: CreateTree([]int{}), sum: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
