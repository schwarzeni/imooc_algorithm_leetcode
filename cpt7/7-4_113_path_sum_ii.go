// https://leetcode-cn.com/problems/path-sum-ii/
// 113. 路径总和 II
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:
//
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]
package cpt6

func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	walkii(root, sum, []int{}, &paths)
	return paths
}

func walkii(node *TreeNode, sum int, path []int, paths *[][]int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	d := sum - node.Val
	if node.Left == nil && node.Right == nil { // 为叶子节点
		if d == 0 {
			*paths = append(*paths, path)
		}
	} else { // 为其它节点
		arr := make([]int, len(path), cap(path))
		copy(arr, path) // 拷贝为一个副本，防止被append修改同一内存处的指针
		walkii(node.Left, d, path, paths)
		walkii(node.Right, d, arr, paths)
	}
}
