// https://leetcode-cn.com/problems/perfect-squares/comments/
// 279 完全平方数
// 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。
// 你需要让组成和的完全平方数的个数最少。
// 输入: n = 12
// 输出: 3
// 解释: 12 = 4 + 4 + 4.
package cpt6

type NodeInfo struct {
	Val  int
	Step int
}

// 题解：使用广度优先遍历
func numSquares(n int) int {
	var (
		queue   []NodeInfo
		visited = make([]bool, n+1)
	)
	for i := 0; i <= n; i++ {
		visited[i] = false
	}
	visited[n] = true
	queue = append(queue, NodeInfo{Val: n, Step: 0})

	for {
		qsize := len(queue)
		if qsize == 0 {
			break
		}
		curr := queue[0]
		queue = queue[1:]

		for i := 1; ; i++ {
			a := curr.Val - i*i
			if a < 0 {
				break
			}
			if a == 0 {
				return curr.Step + 1
			}
			if !visited[a] {
				queue = append(queue, NodeInfo{Val: a, Step: curr.Step + 1})
				visited[a] = true
			}
		}
	}
	return 4
}
