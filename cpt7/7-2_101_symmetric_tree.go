// https://leetcode-cn.com/problems/symmetric-tree/
// 101. 对称二叉树
package cpt6

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return iss(root.Left, root.Right)
}

func iss(ln *TreeNode, rn *TreeNode) bool {
	if ln == nil && rn == nil {
		return true
	}
	if (ln != nil && rn == nil) || (ln == nil && rn != nil) {
		return false
	}
	return ln.Val == rn.Val && iss(ln.Left, rn.Right) && iss(ln.Right, rn.Left)
}
