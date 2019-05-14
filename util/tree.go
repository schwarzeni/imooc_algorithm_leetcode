package util

// 树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL_NODE = -1000000

// 创建二叉树
func CreateTree(arr []int) *TreeNode {
	var (
		size      = len(arr)  // 用于记录传入的数据长度
		nodeQueue []*TreeNode // 队列，用于存放树的子子节点
	)

	if size == 0 { // 此为空树
		return nil
	}

	// 初始化root节点
	root := &TreeNode{Val: arr[0], Left: nil, Right: nil}
	i := 1
	nodeQueue = append(nodeQueue, root)

forloop:
	for i < size {
		queueSize := len(nodeQueue)
		for _, pn := range nodeQueue {
			var (
				ln *TreeNode // 左节点
				rn *TreeNode // 右节点
			)
			if pn == nil { // 父节点为空，则跳过
				continue
			}
			if i >= size {
				break forloop
			}
			if arr[i] != NULL_NODE {
				ln = &TreeNode{Val: arr[i], Left: nil, Right: nil}
			}
			pn.Left = ln
			nodeQueue = append(nodeQueue, ln)
			i++
			if i >= size {
				break forloop
			}
			if arr[i] != NULL_NODE {
				rn = &TreeNode{Val: arr[i], Left: nil, Right: nil}
			}
			pn.Right = rn
			nodeQueue = append(nodeQueue, rn)
			i++
		}
		nodeQueue = nodeQueue[queueSize:]
	}

	return root
}
