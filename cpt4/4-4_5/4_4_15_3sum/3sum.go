// https://leetcode-cn.com/problems/3sum/
package __4_15_3sum

import "sort"

func threeSum(nums []int) [][]int {
	var (
		result [][]int     // 存放结果
		mi     int         // 中idx
		ri     int         // 右idx
		target int         // 中idx和右idx所指向的数字之和
		li     = 0         // 左idx
		size   = len(nums) // 数组的长度
	)

	sort.Ints(nums) // 先排个序

	// 遍历，由于三数之和为0，
	// 所以最小的nums[li]不能大于0，否则就全为正数了
	for li < size-2 && nums[li] <= 0 {
		// 进行一些初始化
		target = -nums[li] // nums[mi] + nums[ri] = - nums[li] = target
		mi = li + 1
		ri = size - 1

		for mi < ri { // 双指针
			curmi := nums[mi]
			curri := nums[ri]
			if curmi+curri == target {
				result = append(result, []int{-target, curmi, curri})

				for mi < ri && nums[mi] == curmi { // 去重，跳过重复的值
					mi++
				}
				for mi < ri && nums[ri] == curri { // 去重，跳过重复的值
					ri--
				}

			} else if curmi+curri < target {
				mi++ // 之和比较小，向左移动mi，增大之和的值
			} else { // curmi+curri > target
				ri-- // 之和比较大，向右移动ri，减小之和的值
			}
		}

		for li < size-2 && -target == nums[li] { // 去重，跳过重复的值
			li++
		}
	}

	return result
}
