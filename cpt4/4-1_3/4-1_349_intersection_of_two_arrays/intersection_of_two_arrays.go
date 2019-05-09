// https://leetcode-cn.com/problems/intersection-of-two-arrays/
// 求两个数组的交集
// 考察set的使用
package intersection_of_two_arrays

import "sort"

// v1 不考虑有序的情况
//func intersection(nums1 []int, nums2 []int) []int {
//	var (
//		set    = make(map[int]bool)
//		result = make([]int, 0, 10)
//	)
//
//	// 将第一个数组中的数据放入set中
//	for _, d := range nums1 {
//		set[d] = false
//	}
//
//	// 再遍历第二个数组，找出相同的数字
//	for _, d := range nums2 {
//		if _, ok := set[d]; ok == true {
//			set[d] = true
//		}
//	}
//
//	for k, v := range set {
//		if v == true {
//			result = append(result, k)
//		}
//	}
//
//	return result
//}
// v2 有序的情况
func intersection(nums1 []int, nums2 []int) []int {
	var (
		n1size  = len(nums1) // 数组1的size
		n2size  = len(nums2) // 数组2的size
		idx1    = 0          // 数组1的移动指针
		idx2    = 0          // 数组2的移动指针
		currNum int
		result  []int
	)

	if n1size == 0 || n2size == 0 {
		return []int{}
	}

	// init
	sort.Ints(nums1)
	sort.Ints(nums2)
	result = make([]int, 0, func() int {
		if n1size > n2size {
			return n2size
		} else {
			return n1size
		}
	}())

	for idx1 < n1size && idx2 < n2size {
		currNum = nums1[idx1]

		for idx1 < n1size && currNum > nums1[idx1] {
			idx1++
		}

		for idx2 < n2size && currNum > nums2[idx2] {
			idx2++
		}

		if idx1 < n1size && idx2 < n2size && currNum == nums1[idx1] && currNum == nums2[idx2] {
			result = append(result, currNum)
		}

		for idx1 < n1size && currNum == nums1[idx1] {
			idx1++
		}

		for idx2 < n2size && currNum == nums2[idx2] {
			idx2++
		}

	}

	return result
}
