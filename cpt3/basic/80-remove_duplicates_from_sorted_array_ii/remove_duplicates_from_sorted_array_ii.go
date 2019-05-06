// 删除数组中出现次数大于2的数字（后面的向前移动）
// 返回数组长度
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/submissions/
package _80_remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	var (
		pi     = 0         // 遍历使用的idx
		ni     = 0         // 赋值使用的idx
		size   = len(nums) // 数组的长度
		pelem  int         // 前一个数组元素的值
		occnum = 0         // 出现次数

	)

	const (
		maxoccur = 2 // 最多出现次数
	)

	for pi < size {
		if pi == 0 { // 初始化
			pelem = nums[pi]
		}
		if nums[pi] == pelem {
			occnum++
			if occnum <= maxoccur {
				nums[ni] = nums[pi]
				ni++
			}
			pi++
		} else { // 上一个元素已经出现结束了，为新一轮计数初始化
			pelem = nums[pi]
			occnum = 0
		}
	}

	return ni
}
