// 0-1背包问题
// 填表
// capacity 5
// id       0   1   2
// weight   1   2   3
// value    6   10  12
//
//    0  1  2  3  4  5
// 0  0  6  6  6  6  6
// 1  0  6  10 16 16 16
// 2  0  6  10 16 18 22
package cpt9

//func knapsack01(v []int, w []int, c int) int {
//	var (
//		vsize = len(v)
//		wsize = len(w)
//		max   = func(d1, d2 int) int {
//			if d1 > d2 {
//				return d1
//			} else {
//				return d2
//			}
//		}
//	)
//	if vsize != wsize || wsize == 0 {
//		return 0
//	}
//
//	var table = make([][]int, vsize)
//	for i := 0; i < vsize; i++ {
//		table[i] = make([]int, c+1)
//	}
//
//	// 填第一行
//	for i := 0; i <= c && w[0] <= i; i++ {
//		table[0][i] = v[0]
//	}
//
//	// 填后续的
//	for i := 1; i < vsize; i++ {
//		for j := 0; j <= c; j++ {
//			table[i][j] = table[i-1][j]
//			if j >= w[i] {
//				table[i][j] = max(table[i][j], v[i]+table[i-1][j-w[i]])
//			}
//		}
//	}
//
//	return table[vsize-1][c]
//}

// 对空间进行优化 2c
//func knapsack01(v []int, w []int, c int) int {
//	var (
//		vsize = len(v)
//		wsize = len(w)
//		max   = func(d1, d2 int) int {
//			if d1 > d2 {
//				return d1
//			} else {
//				return d2
//			}
//		}
//	)
//	if vsize != wsize || wsize == 0 {
//		return 0
//	}
//
//	var table = make([][]int, 2)
//	for i := 0; i < 2; i++ {
//		table[i] = make([]int, c+1)
//	}
//
//	// 填第一行
//	for i := 0; i <= c && w[0] <= i; i++ {
//		table[0][i] = v[0]
//	}
//
//	// 填后续的
//	for i := 1; i < vsize; i++ {
//		for j := 0; j <= c; j++ {
//			table[i%2][j] = table[(i-1)%2][j]
//			if j >= w[i] {
//				table[i%2][j] = max(table[i%2][j], v[i]+table[(i-1)%2][j-w[i]])
//			}
//		}
//	}
//
//	return table[(vsize-1)%2][c]
//}

// 空间优化为 c
func knapsack01(v []int, w []int, c int) int {
	var (
		vsize = len(v)
		wsize = len(w)
		max   = func(d1, d2 int) int {
			if d1 > d2 {
				return d1
			} else {
				return d2
			}
		}
	)
	if vsize != wsize || wsize == 0 {
		return 0
	}

	var table = make([]int, c+1)

	// 填第一行
	for i := 0; i <= c && w[0] <= i; i++ {
		table[i] = v[0]
	}

	// 填后续的
	for i := 1; i < wsize; i++ {
		for j := c; j >= w[i]; j-- {
			table[j] = max(table[j], v[i]+table[j-w[i]])
		}
	}

	return table[c]
}
