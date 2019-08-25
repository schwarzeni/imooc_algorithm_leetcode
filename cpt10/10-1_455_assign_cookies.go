// https://leetcode-cn.com/problems/assign-cookies/
// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
// 对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；
// 并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，
// 这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
//
// 链接：https://leetcode-cn.com/problems/assign-cookies
// 输入: [1,2,3], [1,1]
// 输出: 1
// 解释:
// 你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
// 虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
// 所以你应该输出1。
//
package cpt10

import "sort"

func findContentChildren(g []int, s []int) int {
	var (
		gidx  = len(g) - 1
		sidx  = len(s) - 1
		count = 0
	)

	sort.Ints(g)
	sort.Ints(s)

	for gidx >= 0 && sidx >= 0 {
		if g[gidx] <= s[sidx] {
			sidx -= 1
			count += 1
		}
		gidx -= 1
	}

	return count
}
