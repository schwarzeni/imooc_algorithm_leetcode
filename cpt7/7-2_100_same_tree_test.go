// https://leetcode-cn.com/problems/same-tree/
// 100. 相同的树
package cpt6

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{p: CreateTree([]int{1, 2, 3}), q: CreateTree([]int{1, 2, 3})}, want: true},
		{name: "test2", args: args{p: CreateTree([]int{1, 2, NULL_NODE, 3}), q: CreateTree([]int{1, 2, NULL_NODE, 3})}, want: true},
		{name: "test3", args: args{p: CreateTree([]int{1, 2}), q: CreateTree([]int{1, 2, NULL_NODE, 3})}, want: false},
		{name: "test4", args: args{p: CreateTree([]int{}), q: CreateTree([]int{})}, want: true},
		{name: "test5", args: args{p: CreateTree([]int{1, 2, 5}), q: CreateTree([]int{1, 2, 3})}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
