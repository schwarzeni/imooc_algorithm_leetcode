// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 24. 两两交换链表中的节点
// 示例:
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
package cpt5

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 3, 4})}},
		{name: "test2", args: args{head: CreateLinkList([]int{1})}},
		{name: "test3", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5})}},
		{name: "test4", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 5})}},
		{name: "test5", args: args{head: CreateLinkList([]int{})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
