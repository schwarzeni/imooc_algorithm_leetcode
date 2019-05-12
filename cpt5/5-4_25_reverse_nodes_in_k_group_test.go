// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// 25 k个一组翻转链表
// 给定这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
package cpt5

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 2}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 3}},
		{name: "test3", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 4}},
		{name: "test4", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 5}},
		{name: "test4", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 6}},
		{name: "test5", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5}), k: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); true == true {
				PrintLinkList(got)
			}
		})
	}
}
