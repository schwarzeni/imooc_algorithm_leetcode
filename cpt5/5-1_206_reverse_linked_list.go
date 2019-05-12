// https://leetcode-cn.com/problems/reverse-linked-list/
// 反转链表
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
package cpt5

func reverseList(head *ListNode) *ListNode {
	var (
		currNode = head
		nextNode *ListNode
		prevNode *ListNode
	)

	if head == nil {
		return nil
	}

	for {
		nextNode = currNode.Next
		currNode.Next = prevNode

		prevNode = currNode
		currNode = nextNode

		if currNode == nil {
			break
		}
	}

	return prevNode
}
