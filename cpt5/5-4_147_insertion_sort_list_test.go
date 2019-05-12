// https://leetcode-cn.com/problems/insertion-sort-list/
// 147. 对链表进行插入排序
// 输入: 4->2->1->3
// 输出: 1->2->3->4
package cpt5

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkList([]int{4, 2, 1, 3})}},
		{name: "test2", args: args{head: CreateLinkList([]int{-1, 5, 3, 4, 0})}},
		{name: "test2", args: args{head: CreateLinkList([]int{-1, 0, 1, 2, 5, 4})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
