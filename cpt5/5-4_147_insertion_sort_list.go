// https://leetcode-cn.com/problems/insertion-sort-list/
// 147. 对链表进行插入排序
// 输入: 4->2->1->3
// 输出: 1->2->3->4
package cpt5

func insertionSortList(head *ListNode) *ListNode {
	var (
		prevNode  *ListNode // 当前节点的前一个节点
		currNode  *ListNode // 当前节点
		dummyNode = &ListNode{Next: head}
	)

	if head == nil || head.Next == nil {
		return head
	}
	prevNode = head

	for prevNode.Next != nil {
		if prevNode.Val < prevNode.Next.Val { // 优化：跳过已经有序的序列
			prevNode = prevNode.Next
			continue
		}

		currNode = prevNode.Next
		prevNode.Next = prevNode.Next.Next

		iterNode := dummyNode
		for iterNode != prevNode {
			if iterNode.Next.Val > currNode.Val {
				currNode.Next = iterNode.Next
				iterNode.Next = currNode
				break
			}
			iterNode = iterNode.Next
		}
		if iterNode == prevNode { // 如果该节点还是为最大的，则放回原处
			currNode.Next = prevNode.Next
			prevNode.Next = currNode
			prevNode = prevNode.Next
		}
		// iterNode != prevNode 不用移动prevNode
	}

	return dummyNode.Next
}
