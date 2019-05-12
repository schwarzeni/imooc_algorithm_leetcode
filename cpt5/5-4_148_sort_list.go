// https://leetcode-cn.com/problems/sort-list/
// 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
// 输入: 4->2->1->3
// 输出: 1->2->3->4
package cpt5

func sortList(head *ListNode) *ListNode {
	// 优化？？？使用快慢指针
	count := 0
	for i := head; i != nil; i = i.Next {
		count++
	}
	return mergeSort(head, count)
}

func mergeSort(fstNode *ListNode, size int) *ListNode {
	var (
		secNode = fstNode
	)

	if size < 2 || fstNode == nil || fstNode.Next == nil {
		return fstNode
	}

	for i := 0; i < size/2; i++ {
		secNode = secNode.Next
	}

	pn1Size := size / 2
	pn2Size := size - size/2
	nn1 := mergeSort(fstNode, pn1Size) // 子链表1的结果
	nn2 := mergeSort(secNode, pn2Size) // 子链表2的结果
	dummyHead := &ListNode{Next: nil}
	pn := dummyHead

	for pn1Size > 0 && pn2Size > 0 {
		var cn *ListNode
		if nn1.Val > nn2.Val {
			cn = nn2
			nn2 = nn2.Next // 记录下一个需要指向的节点
			pn2Size--
		} else {
			cn = nn1
			nn1 = nn1.Next // 记录下一个需要指向的节点
			pn1Size--
		}
		pn.Next = cn
		pn = pn.Next
	}

	for pn1Size > 0 {
		cn := nn1
		nn1 = nn1.Next
		pn1Size--
		pn.Next = cn
		pn = pn.Next
	}

	for pn2Size > 0 {
		cn := nn2
		nn2 = nn2.Next
		pn2Size--
		pn.Next = cn
		pn = pn.Next
	}
	pn.Next = nil // 将尾节点置0

	return dummyHead.Next
}
