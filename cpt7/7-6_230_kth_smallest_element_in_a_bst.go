// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 230. 二叉搜索树中第K小的元素
package cpt6

func kthSmallest(root *TreeNode, k int) int {
	res, _ := kths(root, &k)
	return res
}

func kths(root *TreeNode, k *int) (int, bool) {
	if root == nil {
		return 0, false
	}
	if res, ok := kths(root.Left, k); ok {
		return res, ok
	}
	if *k == 1 {
		return root.Val, true
	}
	*k--
	if res, ok := kths(root.Right, k); ok {
		return res, ok
	}
	return 0, false
}
