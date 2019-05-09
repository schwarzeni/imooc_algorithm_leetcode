// https://leetcode-cn.com/problems/3sum-closest/
// 最接近的三数之和
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
package __4_16_3sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var (
		minLen      float64           // 最短的距离
		minTotal    int               // 符合要求的长度
		mi          int               // 中idx
		ri          int               // 右idx
		li          = 0               // 左idx
		curliNum    int               // 当前左idx对于的值
		curmiNum    int               // 当前中idx对于的值
		curriNum    int               // 当前右idx对于的值
		currTotal   int               // 当前三数之和
		currLen     float64           // 当前的距离
		size        = len(nums)       // 数组的长度
		isMinInit   = false           // 最短距离是否初始化
		targetFloat = float64(target) // float64形式的target
	)

	sort.Ints(nums) // 先排个序

	for li < size-2 {
		// 进行一些初始化
		mi = li + 1
		ri = size - 1
		curliNum = nums[li]

		for mi < ri { // 双指针
			curmiNum = nums[mi]
			curriNum = nums[ri]
			currTotal = curliNum + curmiNum + curriNum
			currLen = math.Abs(targetFloat - float64(currTotal))
			if currLen == 0 { // currTotal == target
				return currTotal
			}

			// 判断是否为最小值
			if !isMinInit {
				minLen = currLen
				minTotal = currTotal
				isMinInit = true
			} else if minLen > currLen {
				minLen = currLen
				minTotal = currTotal
			}

			if currTotal < target {
				mi++
			} else {
				ri--
			}

		}

		for li < size-2 && curliNum == nums[li] { // 去重，跳过重复的值
			li++
		}
	}

	return minTotal
}
