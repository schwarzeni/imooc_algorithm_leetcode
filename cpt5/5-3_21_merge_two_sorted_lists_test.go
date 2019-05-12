// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 合并两个有序链表
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
package cpt5

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{l1: CreateLinkList([]int{1, 3, 5, 7, 9}), l2: CreateLinkList([]int{0, 2, 4, 6, 8, 10})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{}), l2: CreateLinkList([]int{0, 2, 4, 6, 8, 10})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{1, 2, 3, 4, 5}), l2: CreateLinkList([]int{6, 7, 8, 9, 10})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); true == true {
				PrintLinkList(got)
			}
		})
	}
}
