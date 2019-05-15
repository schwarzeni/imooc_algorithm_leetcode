// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// 108. 将有序数组转换为二叉搜索树
package cpt6

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l)>>1
	t := &TreeNode{Val: nums[m]}
	t.Left = buildTree(nums, l, m-1)
	t.Right = buildTree(nums, m+1, r)
	return t
}
