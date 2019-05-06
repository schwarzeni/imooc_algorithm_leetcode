// 移除数组中重复的元素，并返回长度
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/submissions/
package _26_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	var (
		pi    = 0         // 遍历使用的idx
		ni    = 0         // 赋值使用的idx
		size  = len(nums) // 数组的长度
		pelem int         // 前一个数组元素的值
	)

	for pi < size {
		if pi == 0 {
			pelem = nums[pi]
		}

		if pelem != nums[pi] {
			nums[ni] = pelem
			ni++
			pelem = nums[pi]
		}

		pi++
	}

	// 处理最后一个元素
	if ni < size {
		nums[ni] = pelem
		ni++
	}

	return ni
}
