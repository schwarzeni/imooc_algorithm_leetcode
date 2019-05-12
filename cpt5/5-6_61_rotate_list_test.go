// https://leetcode-cn.com/problems/rotate-list/
// 61. 旋转链表
// 将链表每个节点向右移动 k 个位置
// 输入: 1->2->3->4->5->NULL, k = 2
// 输出: 4->5->1->2->3->NULL
// 解释:
// 向右旋转 1 步: 5->1->2->3->4->NULL
// 向右旋转 2 步: 4->5->1->2->3->NULL
package cpt5

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
		{name: "test2", args: args{head: CreateLinkList([]int{0, 1, 2}), k: 4}},
		{name: "test3", args: args{head: CreateLinkList([]int{0, 1, 2}), k: 1}},
		{name: "test3", args: args{head: CreateLinkList([]int{0, 1, 2}), k: 0}},
		{name: "test4", args: args{head: CreateLinkList([]int{0}), k: 1}},
		{name: "test5", args: args{head: CreateLinkList([]int{}), k: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); true == true {
				PrintLinkList(got)
			}
		})
	}
}
