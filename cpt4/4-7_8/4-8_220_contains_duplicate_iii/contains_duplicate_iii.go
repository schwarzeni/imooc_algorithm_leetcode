// https://leetcode-cn.com/problems/contains-duplicate-iii/
// 存在重复元素 III
// 给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。
//
// 例 1:
//
// 输入: nums = [1,2,3,1], k = 3, t = 0
// 输出: true
package __8_220_contains_duplicate_iii

// v1 最通用的解法，但是速度慢
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var (
		size = len(nums) // 数组的长度
		li   = 0         // 数组的左idx
		ri   int         // 数组的右idx
	)

	// 不同索引!
	if k == 0 {
		return false
	}

	ri = li + 1
	for li < size-1 {
		// 判断是否相等
		b := nums[ri] - nums[li]
		if b < 0 {
			b = -b
		}
		if b <= t {
			return true
		}

		if ri-li < k && ri < size-1 {
			ri++
			continue
		}

		li++
		// v2 优化？？？？？？？？
		//if t != 0 {
		ri = li + 1
		//}
	}

	return false
}
