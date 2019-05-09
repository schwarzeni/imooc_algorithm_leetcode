package __7_217_contains_duplicate

// v1
func containsDuplicate(nums []int) bool {
	var nmap = make(map[int]int)

	for _, v := range nums {
		nmap[v]++
		if nmap[v] >= 2 {
			return true
		}
	}

	return false
}

//v2 优化内存使用
//func containsDuplicate(nums []int) bool {
//	var (
//		count  = 0
//		curNum int
//	)
//
//	if len(nums) < 1 {
//		return false
//	}
//
//	sort.Ints(nums)
//	curNum = nums[0]
//
//	for _, v := range nums {
//		if v == curNum {
//			count++
//			if count > 1 {
//				return true
//			}
//		} else {
//			count = 1
//			curNum = v
//		}
//	}
//
//	return false
//}
