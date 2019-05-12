// https://leetcode-cn.com/problems/add-two-numbers/
// 两数相加
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
package cpt5

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{l1: CreateLinkList([]int{2, 3}), l2: CreateLinkList([]int{3, 2})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{3, 3, 2}), l2: CreateLinkList([]int{3, 2})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{3, 2}), l2: CreateLinkList([]int{3, 3, 2})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{8, 9, 9}), l2: CreateLinkList([]int{4})}},
		{name: "test1", args: args{l1: CreateLinkList([]int{8, 9, 9}), l2: CreateLinkList([]int{0})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); true == true {
				PrintLinkList(got)
			}
		})
	}
}
