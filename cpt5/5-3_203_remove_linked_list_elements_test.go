// https://leetcode-cn.com/problems/remove-linked-list-elements/
// 移除链表元素
// 输入: 1->2->6->3->4->5->6, val = 6
// 输出: 1->2->3->4->5
package cpt5

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 3, 4, 3, 2, 3}), val: 3}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 1, 1, 1, 1, 1, 2}), val: 1}},
		{name: "test3", args: args{head: CreateLinkList([]int{1, 1, 1, 1, 1, 1}), val: 1}},
		{name: "test4", args: args{head: CreateLinkList([]int{1, 1, 1, 1, 1, 13, 2, 2, 1}), val: 1}},
		{name: "test5", args: args{head: CreateLinkList([]int{}), val: 1}},
		{name: "test6", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), val: 6}},
		{name: "test7", args: args{head: CreateLinkList([]int{1, 1, 2, 2, 1, 1, 3, 3, 1, 1}), val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); true == true {
				PrintLinkList(got)
			}
		})
	}
}
