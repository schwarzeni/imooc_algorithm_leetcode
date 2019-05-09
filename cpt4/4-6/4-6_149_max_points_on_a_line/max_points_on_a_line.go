// https://leetcode-cn.com/problems/max-points-on-a-line/
// 直线上最多的点数
package __6_149_max_points_on_a_line

import (
	"strconv"
)

func maxPoints(points [][]int) int {
	var (
		size     = len(points)
		maxCount = 0
	)

	if size < 3 {
		return size
	}

	for idx, currPoint := range points {
		var (
			kmap      = make(map[string]int)
			sameCount = 0
		)
		for i := idx + 1; i < size; i++ {
			b := points[i][1] - currPoint[1]
			a := points[i][0] - currPoint[0]
			if a != 0 || b != 0 { // 不为同一个点
				g := gcd(a, b)
				//kmap[fmt.Sprintf("%d#%d", a/g, b/g)]++
				kmap[strconv.Itoa(a/g)+strconv.Itoa(b/g)]++
			} else {
				sameCount++
			}
		}
		for _, v := range kmap {
			v = v + sameCount
			if v > maxCount {
				maxCount = v
			}
		}
		if len(kmap) == 0 && sameCount > maxCount {
			maxCount = sameCount
		}
	}

	if maxCount == 0 {
		return 0
	} else {
		return maxCount + 1
	}
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

//func maxPoints(points [][]int) int {
//	type k struct {
//		V1 int // / 结果
//		V2 int // % 结果
//	}
//	var (
//		size     = len(points)
//		maxCount = 0
//	)
//
//	if size < 3 {
//		return size
//	}
//
//	for idx, currPoint := range points {
//		var (
//			kmap      = make(map[k]int) // k 斜率 v 个数
//			sameCount = 0
//		)
//		for i := 0; i < size; i++ {
//			if i != idx {
//				b := points[i][1] - currPoint[1]
//				a := points[i][0] - currPoint[0]
//				if a != 0 || b != 0 { // 不为同一个点
//					currK := k{}
//					if a == 0 { // k = 无穷
//						currK.V2 = 0
//						currK.V1 = math.MaxInt64
//					} else if b == 0 { // k = 0
//						currK.V2 = math.MaxInt64
//						currK.V1 = 0
//					} else {
//						currK.V1 = b / a
//						currK.V2 = b % a
//						// -1 % -2 == -1
//						if b < 0 && a < 0 {
//							currK.V2 = -currK.V2
//							// 1 % -2 == 1
//						} else if b > 0 && a < 0 {
//							currK.V2 = -currK.V2
//						}
//					}
//					kmap[currK]++
//				} else {
//					sameCount++
//				}
//
//			}
//		}
//		for _, v := range kmap {
//			v = v + sameCount
//			if v > maxCount {
//				maxCount = v
//			}
//		}
//		if len(kmap) == 0 && sameCount > maxCount {
//			maxCount = sameCount
//		}
//	}
//
//	if maxCount == 0 {
//		return 0
//	} else {
//		return maxCount + 1
//	}
//}
