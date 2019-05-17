// https://leetcode-cn.com/problems/triangle/
// 120. 三角形最小路径和
// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
// 例如，给定三角形：
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
package cpt9

import "math"

// v1 速度比较慢
//func minimumTotal(triangle [][]int) int {
//	var (
//		walk   func(level int, idx int, counter int)
//		minLen int
//		isInit = false
//	)
//
//	walk = func(level int, idx int, counter int) {
//		if level == len(triangle) {
//			if !isInit {
//				minLen = counter
//				isInit = true
//			} else if minLen > counter {
//				minLen = counter
//			}
//			// 计算长度
//			return
//		}
//
//		walk(level+1, idx, counter+triangle[level][idx])
//		if level > 0 {
//			walk(level+1, idx+1, counter+triangle[level][idx+1])
//		}
//	}
//
//	walk(0, 0, 0)
//
//	return minLen
//}

// v2 记忆搜索
func minimumTotal(triangle [][]int) int {
	var (
		walk  func(level int, idx int) int
		cache [][]int
	)
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for _, y := range triangle {
		a := make([]int, len(y))
		for j, _ := range y {
			a[j] = math.MinInt64
		}
		cache = append(cache, a)
	}

	walk = func(level int, idx int) int {
		if level == len(triangle)-1 {
			return triangle[level][idx]
		}
		var v1, v2 int

		if v1 = cache[level+1][idx]; v1 == math.MinInt64 {
			v1 = walk(level+1, idx)
			cache[level+1][idx] = v1
		}
		if v2 = cache[level+1][idx+1]; v2 == math.MinInt64 {
			v2 = walk(level+1, idx+1)
			cache[level+1][idx+1] = v2
		}
		var v int
		if v1 > v2 {
			v = v2
		} else {
			v = v1
		}
		return v + triangle[level][idx]
	}

	return walk(0, 0)
}

// v3 动态规划

//func minimumTotal(triangle [][]int) int {
//	n:=len(triangle)
//	if n == 0{
//		return 0
//	}
//	for i:=n-2;i>=0;i--{
//		m:=len(triangle[i])
//		for j:=0;j<m;j++{
//			if triangle[i+1][j+1]<triangle[i+1][j] {
//				triangle[i][j]+=triangle[i+1][j+1]
//			}else{
//				triangle[i][j]+=triangle[i+1][j]
//			}
//		}
//	}
//	return triangle[0][0]
//}
