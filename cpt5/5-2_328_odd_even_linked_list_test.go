// https://leetcode-cn.com/problems/odd-even-linked-list/
// 奇偶链表
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
// 请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
// 输入: 2->1->3->5->6->4->7->NULL
// 输出: 2->3->6->7->1->5->4->NULL
package cpt5

import (
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{CreateLinkList([]int{0, 1, 2, 3, 4, 5, 6, 7})}},
		{name: "test2", args: args{CreateLinkList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})}},
		{name: "test3", args: args{CreateLinkList([]int{0})}},
		{name: "test4", args: args{CreateLinkList([]int{})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); true == true {
				PrintLinkList(got)
			}
		})
	}
}
