// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// 83. 删除排序链表中的重复元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
// 示例 1:
//
// 输入: 1->1->2
// 输出: 1->2
package cpt5

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{CreateLinkList([]int{1, 1, 2, 2, 2, 3, 3, 4, 5, 5})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
		{name: "test2", args: args{CreateLinkList([]int{1, 2, 2, 2, 3, 4, 4, 5})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
		{name: "test2", args: args{CreateLinkList([]int{1})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
		{name: "test2", args: args{CreateLinkList([]int{1, 1, 1, 1})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
		{name: "test2", args: args{CreateLinkList([]int{1, 1, 1, 3, 3, 3})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
		{name: "test2", args: args{CreateLinkList([]int{1, 2, 3, 4, 5})}, want: CreateLinkList([]int{1, 2, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
