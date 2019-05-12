// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 24. 两两交换链表中的节点
// 示例:
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
package cpt5

func swapPairs(head *ListNode) *ListNode {
	var (
		dummyHead = &ListNode{Next: head}
		prevNode  = dummyHead
		fstNode   *ListNode // 交换的第一个点
		secNode   *ListNode // 交换的第二个点
	)

	for prevNode.Next != nil && prevNode.Next.Next != nil {
		fstNode = prevNode.Next
		secNode = prevNode.Next.Next

		fstNode.Next = secNode.Next
		secNode.Next = fstNode
		prevNode.Next = secNode

		prevNode = fstNode
	}

	return dummyHead.Next // TODO 内存泄漏？？
}

/** 递归写法
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	list := head.Next
	head.Next = swapPairs(list.Next)
	list.Next = head

	return list
}
*/
