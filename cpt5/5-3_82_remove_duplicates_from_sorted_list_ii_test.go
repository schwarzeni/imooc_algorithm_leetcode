// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
// 82. 删除排序链表中的重复元素 II
// 排序链表
// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
package cpt5

import (
	"testing"
)

func Test_deleteDuplicatesII(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{1, 1, 2, 3, 3, 3})}},
		{name: "test1", args: args{head: CreateLinkList([]int{1, 2, 2, 3, 3, 3, 4})}},
		{name: "test1", args: args{head: CreateLinkList([]int{1, 1, 1, 2, 3})}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 1, 1, 1, 1})}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 2, 3, 4, 4})}},
		{name: "test2", args: args{head: CreateLinkList([]int{1, 1, 3, 5, 6})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesII(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
