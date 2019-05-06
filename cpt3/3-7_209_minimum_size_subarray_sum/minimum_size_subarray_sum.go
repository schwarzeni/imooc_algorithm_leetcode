// 求最短子区间
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/submissions/
package minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {
	var (
		size    = len(nums) // 数组长度
		li      = 0         // 左idx
		ri      = 0         // 右idx
		cursum  = 0         // 当前数之和
		curSize = 0         // 当前的子数组的长度
		minSize = size + 1  // 最小长度
	)

	for {
		if cursum >= s { // 左区间右移
			if li > ri { // 边界条件
				break
			}
			cursum -= nums[li]
			curSize--
			li++
		} else { // 右区间向右移动
			if ri >= size { // 边界条件
				break
			}
			cursum += nums[ri]
			curSize++
			ri++
		}

		if cursum >= s && curSize < minSize {
			minSize = curSize
		}
	}

	if minSize == size+1 {
		return 0
	} else {
		return minSize
	}
}
