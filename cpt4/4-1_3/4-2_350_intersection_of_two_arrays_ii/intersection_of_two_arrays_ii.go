// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// 两个数组的交集 II
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
package intersection_of_two_arrays_ii

import "sort"

// v1 不考虑数组有序的情况
//func intersect(nums1 []int, nums2 []int) []int {
//	var (
//		nmap   = make(map[int]int)
//		result = make([]int, 0, func() int {
//			if len(nums1) > len(nums2) {
//				return len(nums2)
//			} else {
//				return len(nums1)
//			}
//		}())
//	)
//
//	for _, v := range nums1 {
//		nmap[v]++
//	}
//
//	for _, v := range nums2 {
//		if count, ok := nmap[v]; ok == true {
//			if count > 0 {
//				result = append(result, v)
//			}
//			nmap[v]--
//		}
//	}
//
//	return result
//}

// v2 默认数组有序
func intersect(nums1 []int, nums2 []int) []int {
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

		for idx1 < n1size && idx2 < n2size && currNum == nums1[idx1] && currNum == nums2[idx2] {
			result = append(result, currNum)
			idx1++
			idx2++
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
