// https://leetcode-cn.com/problems/palindrome-linked-list/
// 234. 回文链表
// 输入: 1->2->2->1
// 输出: true
package cpt5

func isPalindrome(head *ListNode) bool {
	var (
		ln = head
		rn = head
	)

	if head == nil || head.Next == nil {
		return true
	}

	for rn.Next != nil && rn.Next.Next != nil { // 找到位于中间的节点
		rn = rn.Next.Next
		ln = ln.Next
	}

	rn = ln.Next
	ln.Next = nil
	ln = head
	rn = reverseListNode(rn) // 反转右半部分节点

	for ln != nil && rn != nil { // 进行比较
		if ln.Val != rn.Val {
			return false
		}
		ln = ln.Next
		rn = rn.Next
	}

	return true
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var t *ListNode
	p, q := head, head.Next
	for q != nil {
		t = q.Next
		q.Next = p
		p = q
		q = t
	}
	head.Next = nil
	return p
}
