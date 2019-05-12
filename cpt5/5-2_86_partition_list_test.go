// https://leetcode-cn.com/problems/partition-list/
// 86. 分隔链表
// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
// 你应当保留两个分区中每个节点的初始相对位置。
//
// 示例:
//
// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5
package cpt5

import (
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 4, 3, 2, 5, 2}), x: 3}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 4, 3, 2, 5, 2}), x: 0}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 4, 3, 2, 5, 2}), x: 1}},
		{name: "test2", args: args{head: CreateLinkList([]int{100, 1, 4, 3, 2, 5, 2}), x: 100}},
		{name: "test2", args: args{head: CreateLinkList([]int{4, 3, 2, 5, 2, 1, 1}), x: 2}},
		{name: "test2", args: args{head: CreateLinkList([]int{4, 3, 2, 5, 2, 1, 1}), x: 2}},
		{name: "test2", args: args{head: CreateLinkList([]int{4}), x: 1}},
		{name: "test2", args: args{head: CreateLinkList([]int{4}), x: 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); true == true {
				PrintLinkList(got)
			}
		})
	}
}
