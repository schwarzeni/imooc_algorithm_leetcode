// https://leetcode-cn.com/problems/combination-sum-iii/
// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 示例 2:
// 输入: k = 3, n = 9
// 输出: [[1,2,6], [1,3,5], [2,3,4]]
package cpt8

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	combinationSumHelper3([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, n, k, 0, []int{}, &result)
	if result == nil {
		return [][]int{}
	}
	return result
}

func combinationSumHelper3(candidates []int, target int, size int, startIdx int, sum []int, sums *[][]int) {
	if target < 0 || size < 0 || (target == 0 && size != 0) { // 无结果或者结果长度不符
		return
	}
	if target == 0 {
		*sums = append(*sums, combinationSumCopyArray3(sum))
		return
	}
	asize := len(candidates)
	for i := startIdx; i < asize; i++ {
		if candidates[i] > target {
			break
		}
		combinationSumHelper3(candidates, target-candidates[i], size-1, i+1, append(sum, candidates[i]), sums)
	}
}

func combinationSumCopyArray3(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
