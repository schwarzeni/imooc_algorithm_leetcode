// https://leetcode-cn.com/problems/add-two-numbers/
// 两数相加
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
package cpt5

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l1Node         = l1      // 当前l1节点
		l2Node         = l2      // 当前l2节点
		resultHeadNode *ListNode // 结果链表的head
		prevResultNode *ListNode // 结果链表的上一个node
		carry          = 0       // 进位
	)

	// 23 + 32
	for l1Node != nil && l2Node != nil {
		result := l1Node.Val + l2Node.Val + carry
		carry = result / 10
		rn := &ListNode{Val: result % 10, Next: nil}

		if resultHeadNode == nil {
			resultHeadNode = rn
		} else {
			prevResultNode.Next = rn
		}
		prevResultNode = rn
		l1Node = l1Node.Next
		l2Node = l2Node.Next
	}

	// 233 + 12
	for l1Node != nil {
		result := l1Node.Val + carry
		carry = result / 10
		rn := &ListNode{Val: result % 10, Next: nil}
		if resultHeadNode == nil {
			resultHeadNode = rn
		} else {
			prevResultNode.Next = rn
		}
		prevResultNode = rn
		l1Node = l1Node.Next
	}

	// 12 + 233
	for l2Node != nil {
		result := l2Node.Val + carry
		carry = result / 10
		rn := &ListNode{Val: result % 10, Next: nil}
		if resultHeadNode == nil {
			resultHeadNode = rn
		} else {
			prevResultNode.Next = rn
		}
		prevResultNode = rn
		l2Node = l2Node.Next
	}

	// 98 + 4
	if carry > 0 {
		rn := &ListNode{Val: carry, Next: nil}
		if resultHeadNode == nil {
			resultHeadNode = rn
		} else {
			prevResultNode.Next = rn
		}
		prevResultNode = rn
	}

	return resultHeadNode
}
