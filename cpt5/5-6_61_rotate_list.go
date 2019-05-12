// https://leetcode-cn.com/problems/rotate-list/
// 61. 旋转链表
// 将链表每个节点向右移动 k 个位置
// 输入: 1->2->3->4->5->NULL, k = 2
// 输出: 4->5->1->2->3->NULL
// 解释:
// 向右旋转 1 步: 5->1->2->3->4->NULL
// 向右旋转 2 步: 4->5->1->2->3->NULL
package cpt5

func rotateRight(head *ListNode, k int) *ListNode {
	// 右移k --> 右移 k % len --> 左移 len - k % len
	var (
		hp   = head
		tp   = head
		size = 0 // 链表的大小
	)

	if head == nil || head.Next == nil || k == 0 { // 排除特殊情况
		return head
	}

	for true {
		size++
		if tp.Next == nil { // 到达尾部
			tp.Next = hp // 先将其变成循环链表
			break
		}
		tp = tp.Next
	}

	k = size - k%size
	for k > 0 {
		k--
		tp = tp.Next
	}

	hp = tp.Next  // 设置头指针
	tp.Next = nil // 断开
	return hp
}
