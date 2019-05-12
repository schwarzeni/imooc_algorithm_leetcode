// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 删除链表的倒数第N个节点
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
package cpt5

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		dummyHead = &ListNode{Next: head}
		fastP     = dummyHead
		slowP     = dummyHead
		count     = 0
	)

	// 处理一些特殊情况
	if head == nil || (head.Next == nil) {
		return nil
	}

	for fastP.Next != nil {
		count++
		if count > n {
			slowP = slowP.Next
		}
		fastP = fastP.Next
	}

	slowP.Next = slowP.Next.Next

	return dummyHead.Next
}
