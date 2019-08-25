// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
// 309. 最佳买卖股票时机含冷冻期
// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 示例:
// 输入: [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
package cpt9

func maxProfit(prices []int) int {
	var (
		psize = len(prices)
		max   = func(d1, d2 int) int {
			if d1 > d2 {
				return d1
			} else {
				return d2
			}
		}
		fn func() int
	)

	fn = func() int {
		if psize < 2 {
			return 0
		}
		var (
			sell = make([]int, psize) // 出售的最大
			buy  = make([]int, psize) // 买入的最大
			cool = make([]int, psize) // 冷却的最大
		)
		buy[0] = -prices[0]
		for i := 1; i < psize; i++ {
			sell[i] = buy[i-1] + prices[i]
			buy[i] = max(cool[i-1]-prices[i], buy[i-1])
			cool[i] = max(sell[i-1], cool[i-1])
		}
		return max(sell[psize-1], cool[psize-1])
	}

	return fn()
}
