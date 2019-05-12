// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// 83. 删除排序链表中的重复元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
// 示例 1:
//
// 输入: 1->1->2
// 输出: 1->2
package cpt5

// v1 没注意到是 排序链表，速度比较慢
//func deleteDuplicates(head *ListNode) *ListNode {
//	var (
//		prevNode *ListNode                // 链表的上一个节点
//		currNode = head                   // 链表当前的节点
//		nset     = make(map[int]struct{}) // 用于记录是否重复
//	)
//
//	for currNode != nil {
//		if _, ok := nset[currNode.Val]; ok {
//			if prevNode != nil {
//				prevNode.Next = currNode.Next
//			}
//		} else {
//			nset[currNode.Val] = struct{}{}
//			prevNode = currNode
//		}
//		currNode = currNode.Next
//	}
//
//	return head
//}

//v2
func deleteDuplicates(head *ListNode) *ListNode {
	var (
		prevNode *ListNode // 链表的上一个不同值节点
		currNode = head    // 链表当前的节点
		curVal   int       // 当前的值
	)

	if head == nil {
		return nil
	} else {
		curVal = head.Val
		prevNode = head
	}

	for currNode != nil {
		if currNode.Val != curVal {
			curVal = currNode.Val
			prevNode.Next = currNode
			prevNode = currNode
		}

		if currNode.Next == nil { // 1,2,3,5,5,5
			prevNode.Next = nil
		}
		currNode = currNode.Next
	}

	return head
}
