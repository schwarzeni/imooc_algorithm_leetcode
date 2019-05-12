// https://leetcode-cn.com/problems/reorder-list/
// 143. 重排链表
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
package cpt5

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3})}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4})}},
		{name: "test3", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			PrintLinkList(tt.args.head)
		})
	}
}
