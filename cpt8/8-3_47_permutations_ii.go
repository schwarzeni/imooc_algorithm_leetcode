// https://leetcode-cn.com/problems/permutations-ii/
// 47. 全排列 II
// 输入: [1,1,2]
// 输出:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]
package cpt8

// 别人的解法
// var ret [][]int
//
//func permuteUnique(nums []int) [][]int {
//	if len(nums) == 0 {
//		return ret
//	}
//
//	ret = make([][]int, 0)
//	if len(nums) == 1 {
//		ret = append(ret, nums)
//		return ret
//	}
//
//	sort.Ints(nums)
//	findPermute(nums, 0, len(nums)-1)
//	return ret
//}
//
//func findPermute(nums []int, start int, end int) {
//	if start == end {
//		ret = append(ret, nums)
//		return
//	}
//
//	for i := start; i <= end; i++ {
//		if i != start && nums[i] == nums[start] { // 去重
//			continue
//		}
//
//		nums[start], nums[i] = nums[i], nums[start]
//		newNums := make([]int, end+1)
//		copy(newNums, nums)
//
//		findPermute(newNums, start+1, end)
//	}
//}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	permuteHelperUnique(nums, []int{}, &result)
	return result
}

func permuteHelperUnique(nums []int, currNums []int, result *[][]int) {
	if len(nums) == 1 {
		t := make([]int, len(currNums))
		copy(t, currNums)
		t = append(t, nums[0])

		*result = append(*result, t)
		return
	}

	var nmap = make(map[int]struct{})
	for idx, val := range nums {
		if _, ok := nmap[val]; !ok {
			nmap[val] = struct{}{}
			permuteHelperUnique(removeByIdxUnique(nums, idx), append(currNums, val), result)
		}
	}
}

func removeByIdxUnique(nums []int, idx int) []int {
	var arr []int
	for i, v := range nums {
		if i != idx {
			arr = append(arr, v)
		}
	}
	return arr
}
