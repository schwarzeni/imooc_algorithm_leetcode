// https://leetcode-cn.com/problems/delete-node-in-a-bst/
// 450. 删除二叉搜索树中的节点
package cpt6

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _ = deleteNode(tt.args.root, tt.args.key); true == true {
			}
		})
	}
}
