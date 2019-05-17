// https://leetcode-cn.com/problems/climbing-stairs/
// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 示例 1：
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// 示例 2：
// 输入： 3
// 输出： 3
// 解释： 有三种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶 + 1 阶
// 2.  1 阶 + 2 阶
// 3.  2 阶 + 1 阶
package cpt9

// v1 速度太慢
//func climbStairs(n int) int {
//	var (
//		counter int
//		climb   func(n int)
//		//cache   = make(map[int]struct{})
//	)
//
//	climb = func(step int) {
//		if step > n {
//			return
//		}
//		if n == step {
//			counter++
//			return
//		}
//		climb(step + 1)
//		climb(step + 2)
//	}
//
//	climb(0)
//
//	return counter
//}

// v2 记忆搜索
func climbStairs(n int) int {
	var (
		climb func(n int) int
		cache = make(map[int]int)
	)

	climb = func(step int) int {
		if step == 1 {
			return 1
		}
		if step == 2 {
			return 2
		}
		var (
			v1, v2 int
			ok     bool
		)
		if v1, ok = cache[step-1]; !ok {
			v1 = climb(step - 1)
			cache[step-1] = v1
		}
		if v2, ok = cache[step-2]; !ok {
			v2 = climb(step - 2)
			cache[step-2] = v2
		}
		return v1 + v2
		//return climb(step-1) + climb(step-2)
	}

	return climb(n)
}
