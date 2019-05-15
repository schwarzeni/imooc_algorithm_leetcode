// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
// 129. 求根到叶子节点数字之和
// 输入: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// 输出: 1026
// 解释:
// 从根到叶子节点路径 4->9->5 代表数字 495.
// 从根到叶子节点路径 4->9->1 代表数字 491.
// 从根到叶子节点路径 4->0 代表数字 40.
// 因此，数字总和 = 495 + 491 + 40 = 1026.
package cpt6

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{root: CreateTree([]int{4, 9, 0, 5, 1})}, want: 1026},
		{name: "test2", args: args{root: CreateTree([]int{1, 2, 3})}, want: 25},
		{name: "test3", args: args{root: CreateTree([]int{1})}, want: 1},
		{name: "test4", args: args{root: CreateTree([]int{})}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
