// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
// 注意:
// 可以认为区间的终点总是大于它的起点。
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

// 示例 1:
// 输入: [ [1,2], [2,3], [3,4], [1,3] ]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。

// 示例 2:
// 输入: [ [1,2], [1,2], [1,2] ]
// 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

// 示例 3:
// 输入: [ [1,2], [2,3] ]
// 输出: 0
// 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
// 链接：https://leetcode-cn.com/problems/non-overlapping-intervals
package cpt10

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	return g(intervals)
}

// 贪心算法
func g(intervals [][]int) int {
	var (
		isize   = len(intervals)
		preIdx  = 0
		currIdx = 1
		res     = 1
	)
	if isize == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	for ; currIdx < isize; currIdx++ {
		if intervals[currIdx][0] >= intervals[preIdx][1] {
			res += 1
			preIdx = currIdx
		}
	}

	return isize - res
}

// 动态规划
func dp(intervals [][]int) int {
	var (
		isize      = len(intervals)
		mom        []int
		idx        = 1
		maxSubSize = 1
	)

	if isize <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	mom = make([]int, isize)
	mom[0] = 1

	for idx < isize {
		prevIdx := idx - 1
		for prevIdx >= 0 {
			if intervals[prevIdx][1] <= intervals[idx][0] {
				break
			}
			prevIdx -= 1
		}
		if prevIdx >= 0 {
			mom[idx] += mom[prevIdx] + 1
		} else {
			mom[idx] = 1
		}
		if maxSubSize < mom[idx] {
			maxSubSize = mom[idx]
		}
		idx++
	}

	return isize - maxSubSize
}
