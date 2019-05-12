// https://leetcode-cn.com/problems/odd-even-linked-list/
// 奇偶链表
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
// 请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
// 输入: 2->1->3->5->6->4->7->NULL
// 输出: 2->3->6->7->1->5->4->NULL
package cpt5

func oddEvenList(head *ListNode) *ListNode {
	var (
		prevOddNode  *ListNode
		headOddNode  *ListNode // 奇数节点的头节点
		prevEvenNode *ListNode
		currNode     = head
		isOdd        = false // 是否为奇数节点
	)

	for currNode != nil {
		if isOdd { // 1 3 5 7
			if prevOddNode != nil {
				prevOddNode.Next = currNode
			} else {
				headOddNode = currNode
			}
			prevOddNode = currNode
		} else { // 0 2 4 6
			if prevEvenNode != nil {
				prevEvenNode.Next = currNode
			}
			prevEvenNode = currNode
		}
		currNode = currNode.Next
		isOdd = !isOdd
	}

	if prevEvenNode != nil {
		prevEvenNode.Next = headOddNode
	}
	if prevOddNode != nil {
		prevOddNode.Next = nil
	}

	return head
}
