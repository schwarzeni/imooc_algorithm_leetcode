// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// 111. 二叉树的最小深度
package cpt6

func minDepth(root *TreeNode) int {
	if root == nil { // root为空
		return 0
	}
	return mind(root, 1)
}

func mind(node *TreeNode, level int) int {
	if node == nil {
		return level - 1
	}

	ll := mind(node.Left, level+1)
	rl := mind(node.Right, level+1)

	if ll == level { // 等于当前的level，相当于左子树为空
		return rl
	}
	if rl == level { // 同理
		return ll
	}
	if ll < rl {
		return ll
	} else {
		return rl
	}
}
