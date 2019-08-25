// https://leetcode-cn.com/problems/coin-change/
// 322. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// 示例 1:
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
// 示例 2:
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
// 你可以认为每种硬币的数量是无限的。
package cpt9

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{coins: []int{1, 2, 5}, amount: 11}, want: 3},
		{name: "test2", args: args{coins: []int{2}, amount: 3}, want: -1},
		{name: "test3", args: args{coins: []int{1, 2, 3, 4, 5}, amount: 8}, want: 2},
		{name: "e1超时", args: args{coins: []int{1, 2, 5}, amount: 100}, want: 20},
		{name: "e2超时", args: args{coins: []int{186, 419, 83, 408}, amount: 6249}, want: 20},
		//{name: "test3(可能栈溢出)", args: args{coins: []int{1}, amount: math.MaxInt64}, want: math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
