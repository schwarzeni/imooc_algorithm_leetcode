// https://leetcode-cn.com/problems/combinations/
// 77. 组合
// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// 示例:
// 输入: n = 4, k = 2
// 输出:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]
package cpt8

func combine(n int, k int) [][]int {
	var result [][]int
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	combineHelper(n, k, 1, []int{}, &result)
	if result == nil {
		return [][]int{}
	}
	return result
}

func combineHelper(n int, k int, startIdx int, arr []int, res *[][]int) {
	if len(arr) == k {
		*res = append(*res, combineCopyArray(arr))
		return
	}
	// TODO 优化
	//for i := startIdx; i <= n; i++ {
	for i := startIdx; i <= n-(k-len(arr))+1; i++ {
		combineHelper(n, k, i+1, append(arr, i), res)
	}
}

func combineCopyArray(nums []int) []int {
	newArr := make([]int, len(nums))
	copy(newArr, nums)
	return newArr
}
