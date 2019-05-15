// https://leetcode-cn.com/problems/path-sum/
// 112. 路径总和
// 说明: 叶子节点是指没有子节点的节点。
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \      \
//         7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
package cpt6

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	ln := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return ln == 0
	}
	return hasPathSum(root.Left, ln) || hasPathSum(root.Right, ln)
}
