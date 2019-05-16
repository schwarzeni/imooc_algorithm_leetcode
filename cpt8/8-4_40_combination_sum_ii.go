// https://leetcode-cn.com/problems/combination-sum-ii/
// 40. 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]
package cpt8

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	combinationSumHelper2(candidates, target, []int{}, &result)
	if result == nil {
		return [][]int{}
	}
	return result
}

func combinationSumHelper2(candidates []int, target int, sum []int, sums *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*sums = append(*sums, combinationSumCopyArray2(sum))
		return
	}
	nmap := make(map[int]struct{})
	for idx, v := range candidates {
		if _, ok := nmap[v]; !ok {
			combinationSumHelper2(candidates[idx+1:], target-v, append(sum, v), sums)
			nmap[v] = struct{}{}
		}
	}
}

func combinationSumCopyArray2(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
