// https://leetcode-cn.com/problems/reverse-linked-list-ii/
// 92. 反转链表 II
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
package cpt5

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6}), m: 2, n: 4}, want: nil},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6}), m: 1, n: 6}, want: nil},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6}), m: 1, n: 4}, want: nil},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6}), m: 4, n: 6}, want: nil},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5, 6}), m: 3, n: 3}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); true == true {
				PrintLinkList(got)
			}
		})
	}
}
