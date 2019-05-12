// https://leetcode-cn.com/problems/reorder-list/
// 143. 重排链表
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
package cpt5

// v1 引入数据结构 栈
//func reorderList(head *ListNode) {
//	var (
//		ri  = head
//		li  = head
//		bns []*ListNode
//	)
//
//	if head == nil || head.Next == nil {
//		return
//	}
//
//	// 找到后半部分链表的顶点
//	for ri.Next != nil && ri.Next.Next != nil {
//		ri = ri.Next.Next
//		li = li.Next
//	}
//	li = li.Next
//
//	// 将后半部分元素入栈
//	for li != nil {
//		bns = append(bns, li)
//		li = li.Next
//	}
//
//	li = head
//	for i := len(bns) - 1; i >= 0; i-- {
//		bns[i].Next = li.Next
//		li.Next = bns[i]
//		li = bns[i].Next
//	}
//
//	if bns[0].Next != nil { // 处理尾节点下一个指向的问题
//		li.Next = nil
//	}
//}

// v2 仅在链表上操作
func reorderList(head *ListNode) {
	var (
		ri = head
		li = head
	)

	if head == nil || head.Next == nil {
		return
	}

	// 找到后半部分链表的顶点
	for ri.Next != nil && ri.Next.Next != nil {
		ri = ri.Next.Next
		li = li.Next
	}

	// 反转后半部分节点，同时分割为两个链表
	li.Next = reverListNode(li.Next)
	in := li.Next
	li.Next = nil

	// 将相应的值进行插入
	li = head
	var nli *ListNode
	for {
		nli = li.Next
		nin := in.Next
		in.Next = nli
		li.Next = in

		in = nin
		if in == nil {
			break
		}
		li = li.Next.Next
	}
}

func reverListNode(head *ListNode) *ListNode {
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
