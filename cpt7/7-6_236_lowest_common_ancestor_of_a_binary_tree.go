// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先
package cpt6

// 我想的太复杂了，先贴一段简洁的代码把
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil{
//		return nil
//	}
//	if root == p || root == q{
//		return root
//	}
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//	if left == nil{
//		return right
//	}
//	if right == nil{
//		return left
//	}
//	return root
//}

func lowestCommonAncestorNotBST(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	_, _, r := findLCA(root, p, q)
	return r
}

func findLCA(root *TreeNode, p *TreeNode, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	var (
		ipf = false
		iqf = false
		ipl bool
		iql bool
		ipr bool
		iqr bool
		rnl *TreeNode
		rnr *TreeNode
	)
	//  判断当前节点是否有需要查找的点
	if root.Val == p.Val {
		ipf = true
	} else if root.Val == q.Val {
		iqf = true
	}

	if ipl, iql, rnl = findLCA(root.Left, p, q); true == true {
		if rnl != nil { // 左子树已经有结果了
			return true, true, rnl
		}
		if (ipf && iql) || (ipf && ipl) || (iqf && ipl) || (iqf && iql) { // 左子树和当前节点符合条件
			return true, true, root
		}
	}

	if ipr, iqr, rnr = findLCA(root.Right, p, q); true == true {
		if rnr != nil { // 右子树已经有结果了
			return true, true, rnr
		}
		if (ipf && iqr) || (ipf && ipr) || (iqf && ipr) || (iqf && iqr) { // 右子树和当前节点符合条件
			return true, true, root
		}
	}

	if (ipr && iql) || (ipl && iqr) || (ipl && ipr) || (iql && iqr) { // 当前节点为其祖先节点但是不包含需要查找的点
		return true, true, root
	} else {
		if ipl || iql { // 左子树存在找到的点
			return ipl, iql, nil
		}
		if ipr || iqr { // 右子树存在找到的点
			return ipr, iqr, nil
		}
		return ipf, iqf, nil
	}
}
