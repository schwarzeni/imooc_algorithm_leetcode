// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
// 82. 删除排序链表中的重复元素 II
// 排序链表
// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
package cpt5

func deleteDuplicatesII(head *ListNode) *ListNode {
	var (
		dummyHead  = &ListNode{Next: head} // 初始化时next指向head
		currNode   = dummyHead             // 当前的点的前一个节点
		currVal    int                     // 当前的值
		occurCount = 0                     // 当前的值出现的次数
		prevNode   = dummyHead             // 上一个满足要去的节点
	)

	if head == nil {
		return nil
	}
	currVal = head.Val
	for currNode.Next != nil {
		if currNode.Next.Val == currVal {
			occurCount++
		} else { // currNode.Next.Val != currVal
			if occurCount == 1 {
				prevNode.Next = currNode
				prevNode = currNode
			}
			occurCount = 1
			currVal = currNode.Next.Val
		}
		currNode = currNode.Next
	}
	// 处理最后一个点
	if occurCount > 1 { // 1, 2, 3, 3
		prevNode.Next = nil
	} else { // 1, 2, 2, 3, 3, 4
		prevNode.Next = currNode
	}

	head = dummyHead.Next
	return head
}
