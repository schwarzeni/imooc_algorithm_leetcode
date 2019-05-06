// 求第K大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/comments/
package kth_largest_element_in_an_array

import (
	"sort"
)

// 调用库进行排序
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 快速排序的思想
func findKthLargestII(nums []int, k int) int {
	data, _ := find(nums, k, 0, len(nums)-1)
	return data
}

func find(nums []int, k int, startIdx int, endIdx int) (int, bool) {
	currIdx := partition(nums, startIdx, endIdx)
	if currIdx == len(nums)-k {
		return nums[currIdx], true
	} else if startIdx >= endIdx {
		return 0, false
	}

	if data, isFind := find(nums, k, startIdx, currIdx-1); isFind == true {
		return data, true
	}

	if data, isFind := find(nums, k, currIdx+1, endIdx); isFind == true {
		return data, true
	}

	return 0, false
}

func partition(nums []int, startIdx int, endIdx int) (middleIdx int) {
	var (
		leftIdx   = startIdx
		rightIdx  = endIdx
		middleNum = nums[startIdx]
	)

	if leftIdx == rightIdx {
		middleIdx = leftIdx
		return
	}

	for leftIdx < rightIdx {
		for rightIdx > leftIdx && nums[rightIdx] >= middleNum {
			rightIdx--
		}
		nums[leftIdx] = nums[rightIdx]

		for rightIdx > leftIdx && nums[leftIdx] <= middleNum {
			leftIdx++
		}
		nums[rightIdx] = nums[leftIdx]
	}

	middleIdx = leftIdx
	nums[leftIdx] = middleNum
	return
}
