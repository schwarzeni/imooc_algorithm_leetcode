// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 103. 二叉树的锯齿形层次遍历
// 给定一个二叉树，返回其节点值的锯齿形层次遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
package cpt6

func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		nodeQueue []*TreeNode // 存放节点的queue
		queueSize int         // 记录queue的大小
		result    [][]int     // 记录结果
		fromLtoR  = true      // 从第一行开始
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
		if !fromLtoR {
			levelNodesVal = reverseArr(levelNodesVal)
		}
		fromLtoR = !fromLtoR
		result = append(result, levelNodesVal)
		nodeQueue = nodeQueue[queueSize:] // 更新队列内容
	}

	return result
}

// 反转数组
func reverseArr(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
