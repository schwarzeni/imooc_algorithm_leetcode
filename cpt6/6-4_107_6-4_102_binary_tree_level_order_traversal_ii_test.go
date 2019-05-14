// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 107. 二叉树的层次遍历 II
// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
package cpt6

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
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
			want: [][]int{[]int{15, 7}, []int{9, 20}, []int{3}},
		},
		{
			name: "test2",
			args: args{root: CreateTree([]int{3})},
			want: [][]int{[]int{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
