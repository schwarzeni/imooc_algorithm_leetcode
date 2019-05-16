// https://leetcode-cn.com/problems/subsets-ii/
// 90. 子集 II
// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
// 输入: [1,2,2]
// 输出:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]
package cpt8

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	subsetsWithDupHelper(nums, 0, []int{}, &result)
	result = append(result, []int{})
	return result
}

func subsetsWithDupHelper(nums []int, startIdx int, sum []int, sums *[][]int) {
	var nsize = len(nums)
	if len(sum) > 0 {
		*sums = append(*sums, subsetsWithDupCopyArray(sum))
	}
	if nsize == startIdx { // 返回条件
		return
	}
	var nmap = make(map[int]struct{})
	for i := startIdx; i < nsize; i++ {
		if _, ok := nmap[nums[i]]; !ok {
			nmap[nums[i]] = struct{}{}
			subsetsWithDupHelper(nums, i+1, append(sum, nums[i]), sums)
		}
	}
}

func subsetsWithDupCopyArray(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
