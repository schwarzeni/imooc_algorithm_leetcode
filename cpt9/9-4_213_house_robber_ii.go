//https://leetcode-cn.com/problems/house-robber-ii/
// 213. 打家劫舍 II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
// 示例 1:
// 输入: [2,3,2]
// 输出: 3
// 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 示例 2:
// 输入: [1,2,3,1]
// 输出: 4
// 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//      偷窃到的最高金额 = 1 + 3 = 4 。
package cpt9

func robii(nums []int) int {
	var (
		rb = func(nums []int) int {
			var (
				d1    int
				d2    int
				nsize = len(nums)
			)
			if nsize == 0 {
				return 0
			} else if nsize == 1 {
				return nums[0]
			}
			d1 = nums[0]
			if nums[1] > nums[0] {
				d2 = nums[1]
			} else {
				d2 = nums[0]
			}
			if nsize == 2 {
				return d2
			}
			for i := 2; i < nsize; i++ {
				if nums[i]+d1 > d2 {
					var tmp = d1
					d1 = d2
					d2 = nums[i] + tmp
				} else {
					d1 = d2
				}
			}
			return d2
		}
		nsize = len(nums)
	)
	if nsize == 0 {
		return 0
	} else if nsize == 1 {
		return nums[0]
	}

	m1 := rb(nums[0 : nsize-1])
	m2 := rb(nums[1:nsize])
	if m1 > m2 {
		return m1
	} else {
		return m2
	}
}
