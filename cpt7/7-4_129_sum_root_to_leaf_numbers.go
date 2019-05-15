// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
// 129. 求根到叶子节点数字之和
// 输入: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// 输出: 1026
// 解释:
// 从根到叶子节点路径 4->9->5 代表数字 495.
// 从根到叶子节点路径 4->9->1 代表数字 491.
// 从根到叶子节点路径 4->0 代表数字 40.
// 因此，数字总和 = 495 + 491 + 40 = 1026.
package cpt6

func sumNumbers(root *TreeNode) int {
	sum, _ := walkiii(root)
	return sum
}

func walkiii(node *TreeNode) (sum int, prevLevel []int) {
	if node == nil { // 节点为空
		return
	}
	if node.Left == nil && node.Right == nil { // 为叶子节点
		return node.Val, []int{10}
	}

	ls, ll := walkiii(node.Left)
	rs, rl := walkiii(node.Right)

	for idx, val := range ll {
		ls += node.Val * val
		ll[idx] = val * 10
	}

	for _, val := range rl {
		rs += node.Val * val
		ll = append(ll, val*10)
	}

	return ls + rs, ll
}
