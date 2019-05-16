// https://leetcode-cn.com/problems/permutations/
// 46. 全排列
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]
package cpt8

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	permuteHelper(nums, []int{}, &result)
	return result
}

func permuteHelper(nums []int, currNums []int, result *[][]int) {
	if len(nums) == 1 {
		t := make([]int, len(currNums))
		copy(t, currNums)
		t = append(t, nums[0])
		*result = append(*result, t)
		return
	}

	for idx, val := range nums {
		permuteHelper(removeByIdx(nums, idx), append(currNums, val), result)
	}
}

func removeByIdx(nums []int, idx int) []int {
	var arr []int
	for i, v := range nums {
		if i != idx {
			arr = append(arr, v)
		}
	}
	return arr
}
