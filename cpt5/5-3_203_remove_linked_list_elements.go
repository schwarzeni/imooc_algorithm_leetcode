// https://leetcode-cn.com/problems/remove-linked-list-elements/
// 移除链表元素
// 输入: 1->2->6->3->4->5->6, val = 6
// 输出: 1->2->3->4->5
package cpt5

func removeElements(head *ListNode, val int) *ListNode {
	var (
		prevNode *ListNode
		currNode = head
	)

	for currNode != nil {
		if currNode.Val != val {
			if prevNode != nil {
				prevNode.Next = currNode
			}
			prevNode = currNode
		} else if prevNode == nil { // TODO 可能会内存泄漏
			head = currNode.Next
		}
		currNode = currNode.Next
	}

	// 处理结尾为需要去除的数的情况
	if prevNode != nil {
		prevNode.Next = nil
	}
	return head
}
