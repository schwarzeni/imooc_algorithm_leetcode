// 归并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array
package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		tmpArr = make([]int, m+n, m+n)
		idx1   = 0 // nums1遍历的指针
		idx2   = 0 // nums2遍历的指针
		i      = 0 // tmpArr遍历的指针
	)

	for idx1 < m && idx2 < n {
		if nums1[idx1] < nums2[idx2] {
			tmpArr[i] = nums1[idx1]
			idx1++
		} else {
			tmpArr[i] = nums2[idx2]
			idx2++
		}
		i++
	}

	for idx1 < m {
		tmpArr[i] = nums1[idx1]
		idx1++
		i++
	}

	for idx2 < n {
		tmpArr[i] = nums2[idx2]
		idx2++
		i++
	}

	copy(nums1, tmpArr)
}
