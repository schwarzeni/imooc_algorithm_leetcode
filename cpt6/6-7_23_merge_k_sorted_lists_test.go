// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个排序链表
package cpt6

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{[]*ListNode{
				CreateLinkList([]int{1, 4, 5}),
				CreateLinkList([]int{1, 3, 4}),
				CreateLinkList([]int{2, 6}),
			}},
		},
		{
			name: "test2",
			args: args{[]*ListNode{
				CreateLinkList([]int{1, 4, 5}),
			}},
		},
		{
			name: "e1",
			args: args{[]*ListNode{
				CreateLinkList([]int{}),
				CreateLinkList([]int{}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); true == true {
				PrintLinkList(got)
			}
		})
	}
}
