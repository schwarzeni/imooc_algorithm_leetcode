// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度
package cpt6

func maxDepth(root *TreeNode) int {
	return md(root, 1)
}

func md(node *TreeNode, level int) int {
	if node == nil {
		return level - 1
	}

	ll := md(node.Left, level+1)
	rl := md(node.Right, level+1)
	if ll > rl {
		return ll
	} else {
		return rl
	}
}
