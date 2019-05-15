// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树
package cpt6

func isValidBST(root *TreeNode) bool {
	arr := []int{}
	itraverse(root, &arr)
	size := len(arr)
	if size == 0 || size == 1 {
		return true
	}
	for i := 1; i < size; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func itraverse(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	itraverse(node.Left, arr)
	*arr = append(*arr, node.Val)
	itraverse(node.Right, arr)
}

// 正统思路
//func isValid(root *TreeNode, min, max int) bool {
//	if root==nil {
//		return true
//	}
//	if root.Val<=min || root.Val>=max  {
//		return false
//	}
//
//	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
//}
//
//func isValidBST(root *TreeNode) bool {
//	return isValid(root, math.MinInt64, math.MaxInt64)
//}
