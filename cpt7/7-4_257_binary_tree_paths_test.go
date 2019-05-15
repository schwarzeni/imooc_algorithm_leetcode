// https://leetcode-cn.com/problems/binary-tree-paths/
// 257. 二叉树的所有路径
// 输入:
//
//    1
//  /   \
// 2     3
//  \
//   5
//
// 输出: ["1->2->5", "1->3"]
//
// 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
package cpt6

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{root: CreateTree([]int{1, 2, 3, NULL_NODE, 5})}, want: []string{"1->2->5", "1->3"}},
		{name: "test2", args: args{root: CreateTree([]int{1})}, want: []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
