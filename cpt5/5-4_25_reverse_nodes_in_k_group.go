// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// 25 k个一组翻转链表
// 给定这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
package cpt5

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		dummyHead = &ListNode{Next: head}
		prevNode  = dummyHead
		count     = 0
		kPrevNode = dummyHead // k区间之间的节点
		kNextNode *ListNode   // k区间之后的节点
	)

	if k < 2 || head == nil {
		return head
	}

	for prevNode.Next != nil {
		count++
		if count == k {
			kNextNode = prevNode.Next.Next
			kPrevNode = swapNodes(kPrevNode, kNextNode)
			count = 0
			prevNode = kPrevNode
		} else {
			prevNode = prevNode.Next
		}
	}

	return dummyHead.Next
}

// 交换节点
func swapNodes(prevNode *ListNode, nextNode *ListNode) *ListNode {
	var (
		pn = prevNode.Next
		cn = prevNode.Next.Next
		nn = prevNode.Next.Next.Next
	)

	for cn != nextNode {
		nn = cn.Next
		cn.Next = pn
		pn = cn
		cn = nn
	}

	prevNode.Next.Next = nextNode
	tn := prevNode.Next
	prevNode.Next = pn
	return tn
}
