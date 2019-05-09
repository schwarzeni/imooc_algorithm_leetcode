// https://leetcode-cn.com/problems/contains-duplicate-ii/
// 存在重复元素 II
// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
// 示例 1:
// 入: nums = [1,2,3,1], k = 3
// 出: true
package __7_219_contains_duplicate_ii

//v1 速度比较慢
//func containsNearbyDuplicate(nums []int, k int) bool {
//	var (
//		size = len(nums) // 数组的长度
//		li   = 0         // 数组的左idx
//		ri   int         // 数组的右idx
//	)
//
//	for ; li < size-1; li++ {
//		ri = li + 1
//		for ; ri < size && ri <= li+k; ri++ {
//			if nums[ri] == nums[li] {
//				return true
//			}
//
//		}
//	}
//
//	return false
//}

// v2 使用查找表以空间换时间提高速度
//func containsNearbyDuplicate(nums []int, k int) bool {
//	var (
//		nmap = make(map[int][2]int) // idx=0 prev_idx idx=1 prev_size
//		size = len(nums)
//	)
//
//	// 1 , 1
//	// 边界
//	if size < 2 || k < 1 {
//		return false
//	}
//
//	for idx, val := range nums {
//		if v, ok := nmap[val]; ok {
//			b := idx - v[0]
//			if b <= k {
//				return true
//			}
//			if b < v[1] || v[1] == -1 {
//				v[1] = b
//			}
//			v[0] = idx
//			nmap[val] = v
//		} else {
//			// idx 初始化为-1
//			nmap[val] = [2]int{idx, -1}
//		}
//	}
//
//	// 优化：直接在上面判断就可以了
//	//for _, v := range nmap {
//	//	if v[1] <= k && v[1] > 0 {
//	//		return true
//	//	}
//	//}
//
//	return false
//}

// v3 优化内存使用，精简代码
func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		nmap   = make(map[int]int)
		preIdx int  // 同一个数上一个idx
		hasIdx bool // 是否已经存在
	)

	for idx, val := range nums {
		if preIdx, hasIdx = nmap[val]; hasIdx { // 取最近的一次
			if idx-preIdx <= k {
				return true
			}
		}
		nmap[val] = idx
	}

	return false
}
