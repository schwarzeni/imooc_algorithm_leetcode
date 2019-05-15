// https://leetcode-cn.com/problems/path-sum-iii/
// 437. 路径总和 III
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//       10
//      /  \
//     5   -3
//    / \    \
//   3   2   11
//  / \   \
// 3  -2   1
//
// 返回 3。和等于 8 的路径有:
//
// 1.  5 -> 3
// 2.  5 -> 2 -> 1
// 3.  -3 -> 11
package cpt6

import "testing"

func Test_pathSumIII(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{root: CreateTree([]int{10, 5, -3, 3, 2, NULL_NODE, 11, 3, -2, NULL_NODE, 1}), sum: 8},
			want: 3,
		},
		{
			name: "e1",
			args: args{root: CreateTree([]int{1, NULL_NODE, 2, NULL_NODE, 3, NULL_NODE, 4, NULL_NODE, 5}), sum: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSumIII(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSumIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
