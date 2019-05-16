// https://leetcode-cn.com/problems/combination-sum-iii/
// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 示例 2:
// 输入: k = 3, n = 9
// 输出: [[1,2,6], [1,3,5], [2,3,4]]
package cpt8

import (
	"fmt"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{name: "test1", args: args{k: 3, n: 9}},
		{name: "test2", args: args{k: 3, n: 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); true == true {
				fmt.Println(got)
			}
		})
	}
}
