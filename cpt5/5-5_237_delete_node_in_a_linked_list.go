// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
// 237. 删除链表中的节点
// 链表至少包含两个节点
// 链表中所有节点的值都是唯一的
// 给定的节点为非末尾节点并且一定是链表中的一个有效节点
// 输入: head = [4,5,1,9], node = 5
// 输出: [4,1,9]
package cpt5

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
