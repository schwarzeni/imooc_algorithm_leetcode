// https://leetcode-cn.com/problems/combination-sum/
// 39. 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//   [7],
//   [2,2,3]
// ]
// 示例 2:
// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
package cpt8

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	combinationSumHelper(candidates, target, []int{}, &result)
	if result == nil {
		return [][]int{}
	}
	return result
}

func combinationSumHelper(candidates []int, target int, sum []int, sums *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*sums = append(*sums, combinationSumCopyArray(sum))
		return
	}
	for idx, v := range candidates {
		combinationSumHelper(candidates[idx:], target-v, append(sum, v), sums)
	}
}

func combinationSumCopyArray(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
