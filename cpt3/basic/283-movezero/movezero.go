// 将0移动到末尾
// https://leetcode-cn.com/problems/move-zeroes/
package _283_movezero

func moveZeroes(nums []int) {
	movezerosS2(nums)
}

// 思路:
// 将非零元素移至前方
// 最后将剩余元素制0
func movezerosS1(nums []int) {
	var (
		pi   = 0         // 遍历使用的idx
		ni   = 0         // 赋值使用的idx
		size = len(nums) // 数组的长度
	)

	for pi < size {
		if nums[pi] != 0 {
			nums[ni] = nums[pi]
			ni++
		}
		pi++
	}

	for ni < size {
		nums[ni] = 0
		ni++
	}
}

// 思路:
// 将非零元素和零元素交换位置
func movezerosS2(nums []int) {
	var (
		pi   = 0         // 遍历使用的idx
		ni   = 0         // 赋值使用的idx
		size = len(nums) // 数组的长度
	)
	for pi < size {
		if nums[pi] != 0 {
			if ni != pi { // 全为非零元素的情况
				nums[ni], nums[pi] = nums[pi], nums[ni]
			}
			ni++
		}
		pi++
	}
}
