// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 102. 二叉树的层次遍历
// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
package cpt6

func levelOrder(root *TreeNode) [][]int {
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
		result = append(result, levelNodesVal)
		nodeQueue = nodeQueue[queueSize:] // 更新队列内容
	}

	return result
}
