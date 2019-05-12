// https://leetcode-cn.com/problems/reverse-linked-list/
// 反转链表
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
package cpt5

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6})}, want: nil},
		{name: "test2", args: args{head: CreateLinkList([]int{1})}, want: nil},
		{name: "test3", args: args{head: CreateLinkList([]int{})}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
