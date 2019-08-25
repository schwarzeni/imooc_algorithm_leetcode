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

// 带缓存的深度优先
//func coinChange(coins []int, amount int) int {
//	var (
//		dfs   func(t int) int
//		cache = make(map[int]int)
//	)
//
//	dfs = func(t int) int {
//		var (
//			minSize = -1
//		)
//		if t < 0 {
//			return minSize
//		}
//		if t == 0 {
//			return 0
//		}
//		if v, ok := cache[t]; ok {
//			return v
//		}
//		for _, v := range coins {
//			res := dfs(t - v)
//			if res == -1 {
//				continue
//			} else if minSize == -1 || minSize > res {
//				minSize = res + 1
//			}
//		}
//		cache[t] = minSize
//		return minSize
//	}
//
//	return dfs(amount)
//}

// 动态规划
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = 999999
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] != 999999 {
		return dp[amount]
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
