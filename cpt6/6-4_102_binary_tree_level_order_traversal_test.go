// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 102. 二叉树的层次遍历
// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
package cpt6

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{root: CreateTree([]int{3, 9, 20, NULL_NODE, NULL_NODE, 15, 7})},
			want: [][]int{[]int{3}, []int{9, 20}, []int{15, 7}},
		},
		{
			name: "test2",
			args: args{root: CreateTree([]int{3})},
			want: [][]int{[]int{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
