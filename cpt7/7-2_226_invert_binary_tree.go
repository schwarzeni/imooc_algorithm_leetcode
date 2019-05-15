// https://leetcode-cn.com/problems/invert-binary-tree/
// 226. 翻转二叉树
package cpt6

func invertTree(root *TreeNode) *TreeNode {
	it(root)
	return root
}

func it(node *TreeNode) {
	if node == nil {
		return
	}
	it(node.Left)
	it(node.Right)
	node.Left, node.Right = node.Right, node.Left
}
