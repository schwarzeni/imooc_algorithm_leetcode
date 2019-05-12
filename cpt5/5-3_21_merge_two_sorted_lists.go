// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 合并两个有序链表
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
package cpt5

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead = &ListNode{Next: nil}
		prevNode  = dummyHead
		l1Node    = l1
		l2Node    = l2
	)

	for l1Node != nil && l2Node != nil {
		if l1Node.Val > l2Node.Val {
			prevNode.Next = l2Node
			l2Node = l2Node.Next
		} else {
			prevNode.Next = l1Node
			l1Node = l1Node.Next
		}
		prevNode = prevNode.Next
	}

	for l1Node != nil {
		prevNode.Next = l1Node
		prevNode = prevNode.Next
		l1Node = l1Node.Next
	}

	for l2Node != nil {
		prevNode.Next = l2Node
		prevNode = prevNode.Next
		l2Node = l2Node.Next
	}

	return dummyHead.Next // TODO 可能内存泄漏??
}
