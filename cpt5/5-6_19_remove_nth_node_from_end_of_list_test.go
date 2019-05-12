// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 删除链表的倒数第N个节点
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
package cpt5

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), n: 1}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), n: 5}},
		{name: "test3", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), n: 2}},
		{name: "test4", args: args{head: CreateLinkList([]int{1}), n: 1}},
		{name: "test5", args: args{head: CreateLinkList([]int{1}), n: 4}},
		{name: "test6", args: args{head: CreateLinkList([]int{}), n: 1}},
		{name: "test7", args: args{head: CreateLinkList([]int{1, 2}), n: 1}},
		{name: "test8", args: args{head: CreateLinkList([]int{1, 2}), n: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); true == true {
				PrintLinkList(got)
			}
		})
	}
}
