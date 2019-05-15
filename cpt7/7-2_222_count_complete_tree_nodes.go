// https://leetcode-cn.com/problems/count-complete-tree-nodes/
// 222. 完全二叉树的节点个数
// 另一种思路
// https://leetcode-cn.com/problems/count-complete-tree-nodes/comments/3384
// https://leetcode-cn.com/problems/count-complete-tree-nodes/comments/20211
package cpt6

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
