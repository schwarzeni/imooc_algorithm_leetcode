// https://leetcode-cn.com/problems/subsets/
// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 示例:
// 输入: nums = [1,2,3]
// 输出:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]
package cpt8

func subsets(nums []int) [][]int {
	var result [][]int
	subsetsHelper(nums, 0, []int{}, &result)
	result = append(result, []int{})
	return result
}

func subsetsHelper(nums []int, startIdx int, sum []int, sums *[][]int) {
	var nsize = len(nums)
	if len(sum) > 0 {
		*sums = append(*sums, subsetsHelperCopyArray(sum))
	}
	if nsize == startIdx { // 返回条件
		return
	}
	for i := startIdx; i < nsize; i++ {
		subsetsHelper(nums, i+1, append(sum, nums[i]), sums)
	}
}

func subsetsHelperCopyArray(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
