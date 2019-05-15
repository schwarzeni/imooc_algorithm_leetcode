// https://leetcode-cn.com/problems/balanced-binary-tree/
// 110. 平衡二叉树
package cpt6

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	d := getDepth(root.Left, 1) - getDepth(root.Right, 1)
	return d >= -1 && d <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(node *TreeNode, level int) int {
	if node == nil {
		return level - 1
	}
	ll := getDepth(node.Left, level+1)
	rl := getDepth(node.Right, level+1)
	if ll > rl {
		return ll
	} else {
		return rl
	}
}
