// https://leetcode-cn.com/problems/partition-list/
// 86. 分隔链表
// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
// 你应当保留两个分区中每个节点的初始相对位置。
//
// 示例:
//
// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5
package cpt5

func partition(head *ListNode, x int) *ListNode {
	var (
		currNode  = head    // 当前的元素
		lTailNode *ListNode // 小于x1的元素的尾节点
		rHeadNode *ListNode // 大于x的元素的子链表的头元素
		rPrevNode *ListNode // 大于x1的元素的子链表的上一个元素
	)

	for currNode != nil {
		if currNode.Val < x {
			if lTailNode != nil {
				lTailNode.Next = currNode
			} else {
				head = currNode // 左链表的第一个节点设为head
			}
			lTailNode = currNode
		} else {
			if rHeadNode == nil {
				rHeadNode = currNode
			} else {
				rPrevNode.Next = currNode
			}
			rPrevNode = currNode
		}
		currNode = currNode.Next
	}

	if lTailNode != nil {
		lTailNode.Next = rHeadNode
	}
	if rPrevNode != nil {
		rPrevNode.Next = nil
	}

	return head
}
