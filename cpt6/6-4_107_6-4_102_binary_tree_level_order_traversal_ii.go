// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 107. 二叉树的层次遍历 II
// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
package cpt6

func levelOrderBottom(root *TreeNode) [][]int {
	var (
		nodeQueue []*TreeNode // 存放节点的queue
		queueSize int         // 记录queue的大小
		result    [][]int     // 记录结果
	)

	if root == nil { // 排除掉root为空的情况
		return result
	}

	nodeQueue = append(nodeQueue, root) // 将root节点加入队列
	for {
		var levelNodesVal []int
		queueSize = len(nodeQueue)
		if queueSize == 0 {
			break
		}
		for _, n := range nodeQueue {
			levelNodesVal = append(levelNodesVal, n.Val)
			if n.Left != nil {
				nodeQueue = append(nodeQueue, n.Left)
			}
			if n.Right != nil {
				nodeQueue = append(nodeQueue, n.Right)
			}
		}
		// 头插
		//result = append([][]int{levelNodesVal}, result...)
		result = append(result, levelNodesVal)
		nodeQueue = nodeQueue[queueSize:] // 更新队列内容
	}

	// 反转数组
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
