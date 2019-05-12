// https://leetcode-cn.com/problems/reverse-linked-list-ii/
// 92. 反转链表 II
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
package cpt5

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var (
		currNode = head
		nextNode *ListNode
		prevNode *ListNode
		pNode    *ListNode // 需要反转的子链表头节点的前一个节点
		idx      = 1
	)

	if head == nil {
		return nil
	}

	if m == n {
		return head
	}

	for {
		nextNode = currNode.Next
		if idx == m-1 { // 注意，m==1时pNode为nil
			pNode = currNode
		}

		if idx >= m && idx <= n { // 需要反转的节点
			currNode.Next = prevNode
			prevNode = currNode
		}

		if idx == n { // 为需要反转的最后一个节点
			if pNode != nil {
				pNode.Next.Next = nextNode
				pNode.Next = currNode
			} else {
				head.Next = nextNode
				head = currNode
			}
		}

		idx++
		currNode = nextNode
		if currNode == nil {
			break
		}
	}

	return head
}
