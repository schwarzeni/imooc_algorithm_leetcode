// https://leetcode-cn.com/problems/delete-node-in-a-bst/
// 450. 删除二叉搜索树中的节点
package cpt6

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	root = deleteNodeFunc(root, key)
	return root
}

func deleteNodeFunc(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key { // 删除操作
		if root.Left == nil && root.Right == nil { // 为孩子节点
			root = nil
			return root
		}
		if root.Left == nil { // 为只有一侧
			root = root.Right
			return root
		}
		if root.Right == nil {
			root = root.Left
			return root
		}

		mostLeftNode := root.Right
		for {
			// 一直遍历找到最左节点
			if mostLeftNode != nil && mostLeftNode.Left != nil {
				mostLeftNode = mostLeftNode.Left
			} else {
				break
			}
		}
		root.Val = mostLeftNode.Val
		root.Right = deleteNodeFunc(root.Right, root.Val)
	} else if root.Val < key {
		root.Right = deleteNodeFunc(root.Right, key)
	} else {
		root.Left = deleteNodeFunc(root.Left, key)
	}
	return root
}
