// 数组中仅有 0，1，2 三个元素，请对其进行排序
// https://leetcode-cn.com/problems/sort-colors/submissions/
package _80_sort_colors

func sortColors(nums []int) {
	var (
		count_0 = 0         // 元素0的个数
		count_1 = 0         // 元素1的个数
		count_2 = 0         // 元素2的个数
		size    = len(nums) // 数组长度
	)

	for i := 0; i < size; i++ {
		switch nums[i] {
		case 0:
			count_0++
		case 1:
			count_1++
		case 2:
			count_2++
		}
	}

	for i := 0; i < size; i++ {
		if count_0 > 0 {
			nums[i] = 0
			count_0--
			continue
		}
		if count_1 > 0 {
			nums[i] = 1
			count_1--
			continue
		}
		if count_2 > 0 {
			nums[i] = 2
			count_2--
			continue
		}
	}
}

// 三路快排
// 对撞指针
func sortColorsBetter(nums []int) {
	var (
		size    = len(nums) // 第二个idx
		fstTail = -1        // 第一个元素的尾部idx
		thHead  = size      // 第三个元素的首部idx
		i       = 0         // 遍历使用的idx
	)

	for fstTail < thHead && i < thHead {
		switch nums[i] {
		case 0:
			fstTail++
			nums[i], nums[fstTail] = nums[fstTail], nums[i]
			i++
		case 1:
			i++
		case 2:
			thHead--
			nums[i], nums[thHead] = nums[thHead], nums[i]
		}
	}
}
