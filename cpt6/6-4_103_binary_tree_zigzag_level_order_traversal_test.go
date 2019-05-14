// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 103. 二叉树的锯齿形层次遍历
// 给定一个二叉树，返回其节点值的锯齿形层次遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
package cpt6

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
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
			want: [][]int{[]int{3}, []int{20, 9}, []int{15, 7}},
		},
		//{
		//	name: "test2",
		//	args: args{root: CreateTree([]int{3})},
		//	want: [][]int{[]int{3}},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
