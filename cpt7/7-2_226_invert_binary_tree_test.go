// https://leetcode-cn.com/problems/invert-binary-tree/
// 226. 翻转二叉树
package cpt6

import (
	"fmt"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		//{name: "test1", args: args{CreateTree([]int{4, 2, 7, 1, 3, 6, 9})}},
		//      4
		//   /   \
		//  2     7
		//   \   /
		//    3 6
		//      4
		//   /   \
		//  7     2
		//   \   /
		//    6 3
		{name: "test2", args: args{CreateTree([]int{4, 2, 7, NULL_NODE, 3, 6, NULL_NODE})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); true == true {
				fmt.Println(got)
			}
		})
	}
}
