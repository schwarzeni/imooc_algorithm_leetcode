// 将不要的元素移至末尾
// 并返回剩余元素长度
// https://leetcode-cn.com/problems/remove-element/
package _27_removeelement

func removeElement(nums []int, val int) int {
	var (
		pi   = 0         // 遍历使用的idx
		ni   = 0         // 赋值使用的idx
		size = len(nums) // 数组的长度
	)

	for pi < size {
		if nums[pi] != val {
			nums[ni] = nums[pi]
			ni++
		}
		pi++
	}

	return ni
}
