// https://leetcode-cn.com/problems/binary-tree-right-side-view/
// 199. 二叉树的右视图
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
package cpt6

func rightSideView(root *TreeNode) []int {
	var (
		nodeQueue []*TreeNode // 存放节点的queue
		queueSize int         // 记录queue的大小
		result    []int       // 记录结果
	)

	if root == nil { // 排除掉root为空的情况
		return result
	}

	nodeQueue = append(nodeQueue, root) // 将root节点加入队列
	for {
		var lastVal int // 当前层最后一个节点的值
		queueSize = len(nodeQueue)
		if queueSize == 0 {
			break
		}
		for _, n := range nodeQueue {
			lastVal = n.Val
			if n.Left != nil {
				nodeQueue = append(nodeQueue, n.Left)
			}
			if n.Right != nil {
				nodeQueue = append(nodeQueue, n.Right)
			}
		}
		result = append(result, lastVal)
		nodeQueue = nodeQueue[queueSize:] // 更新队列内容
	}
	return result
}
