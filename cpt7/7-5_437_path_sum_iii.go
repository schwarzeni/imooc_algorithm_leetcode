// https://leetcode-cn.com/problems/path-sum-iii/
// 437. 路径总和 III
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//       10
//      /  \
//     5   -3
//    / \    \
//   3   2   11
//  / \   \
// 3  -2   1
//
// 返回 3。和等于 8 的路径有:
//
// 1.  5 -> 3
// 2.  5 -> 2 -> 1
// 3.  -3 -> 11
package cpt6

func pathSumIII(root *TreeNode, sum int) int {
	//return pathSumiii(root, sum, sum)
	var count = 0
	var dfs func(root *TreeNode, num int)
	var s func(root *TreeNode, num int)
	s = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		num -= root.Val
		if num == 0 {
			count++
		}
		s(root.Left, num)
		s(root.Right, num)
	}
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		s(root, num)
		dfs(root.Left, num)
		dfs(root.Right, num)
	}
	dfs(root, sum)

	return count
}

//var hmap = make(map[*TreeNode]struct{}) // 记录以及作为head使用的过的点
//func pathSumiii(node *TreeNode, currSum int, sum int) (count int) {
//	if node == nil {
//		return
//	}
//	d := currSum - node.Val // 剩下的值
//
//	// 身份为结束点
//	if d == 0 {
//		count++
//	}
//	// 1. 身份为起始点 2. 身份为中间点
//	if node.Left != nil {
//		if _, ok := hmap[node.Left]; !ok {
//			count += pathSumiii(node.Left, sum, sum)
//			hmap[node.Left] = struct{}{}
//		}
//		count += pathSumiii(node.Left, d, sum)
//	}
//	if node.Right != nil {
//		if _, ok := hmap[node.Right]; !ok {
//			count += pathSumiii(node.Right, sum, sum)
//			hmap[node.Right] = struct{}{}
//		}
//		count += pathSumiii(node.Right, d, sum)
//	}
//	return
//}
