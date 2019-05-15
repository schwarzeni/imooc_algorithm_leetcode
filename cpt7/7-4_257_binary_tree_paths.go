// https://leetcode-cn.com/problems/binary-tree-paths/
// 257. 二叉树的所有路径
// 输入:
//
//    1
//  /   \
// 2     3
//  \
//   5
//
// 输出: ["1->2->5", "1->3"]
//
// 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
package cpt6

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var paths []string // 用于存放path的slice
	walk(root, "", &paths)
	return paths
}

func walk(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}
	path += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path)
	} else {
		path += "->"
		walk(node.Left, path, paths)
		walk(node.Right, path, paths)
	}
}
