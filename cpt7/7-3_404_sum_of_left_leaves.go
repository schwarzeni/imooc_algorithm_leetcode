// https://leetcode-cn.com/problems/sum-of-left-leaves/
// 计算给定二叉树的所有左叶子之和。
package cpt6

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil { // 叶子节点，不做统计
		return 0
	}
	count := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + count
	} else {
		return count
	}
}
