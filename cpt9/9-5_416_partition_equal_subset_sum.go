// https://leetcode-cn.com/problems/partition-equal-subset-sum/
// 416. 分割等和子集
// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//注意:
//每个数组中的元素不会超过 100
//数组的大小不会超过 200
//示例 1:
//输入: [1, 5, 11, 5]
//输出: true
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//示例 2:
//输入: [1, 2, 3, 5]
//输出: false
//解释: 数组不能分割成两个元素和相等的子集.
package cpt9

// 超级慢的解法
//func canPartition(nums []int) bool {
//	var (
//		datamap = make(map[int]struct{})
//		nsize   = len(nums)
//		total   int // 所有数之和
//		target  int // 需要求的数字
//	)
//	for _, v := range nums {
//		total += v
//	}
//	if nsize < 2 || total%2 != 0 {
//		return false
//	}
//	target = total / 2
//	for i := 0; i < nsize; i++ {
//		var newData []int
//		for k, _ := range datamap {
//			if k+nums[i] == target {
//				if total > 0 {
//					total = -1
//				} else {
//					return true
//				}
//			}
//			if _, ok := datamap[nums[i]+k]; !ok {
//				newData = append(newData, nums[i]+k)
//			}
//		}
//		for _, v := range newData {
//			datamap[v] = struct{}{}
//		}
//		if nums[i] == target {
//			if total > 0 {
//				total = -1
//			} else {
//				return true
//			}
//		}
//		if _, ok := datamap[nums[i]]; !ok {
//			datamap[nums[i]] = struct{}{}
//		}
//	}
//	return false
//}

// 背包解法
func canPartition(nums []int) bool {
	var (
		nsize    = len(nums)
		total    int
		capacity int
		cache    []bool
	)
	for _, v := range nums {
		total += v
	}
	if nsize < 2 || total%2 != 0 {
		return false
	}
	capacity = total / 2
	cache = make([]bool, capacity+1)
	for i := 0; i <= capacity; i++ {
		cache[i] = nums[0] == i
	}

	for i := 1; i < nsize; i++ {
		for j := capacity; j >= nums[i]; j-- {
			cache[j] = cache[j] || cache[j-nums[i]]
		}
	}

	return cache[capacity]
}
