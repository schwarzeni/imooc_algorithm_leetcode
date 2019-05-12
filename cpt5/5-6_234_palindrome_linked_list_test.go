// https://leetcode-cn.com/problems/palindrome-linked-list/
// 234. 回文链表
// 输入: 1->2->2->1
// 输出: true
package cpt5

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{CreateLinkList([]int{1, 2, 3, 2, 1})}, want: true},
		{name: "test2", args: args{CreateLinkList([]int{1, 2, 2, 1})}, want: true},
		{name: "test3", args: args{CreateLinkList([]int{1, 2, 3, 1})}, want: false},
		{name: "test4", args: args{CreateLinkList([]int{})}, want: true},
		{name: "test5", args: args{CreateLinkList([]int{1})}, want: true},
		{name: "test6", args: args{CreateLinkList([]int{1, 0, 0})}, want: false},
		{name: "test7", args: args{CreateLinkList([]int{1, 0, 1})}, want: true},
		{name: "test8", args: args{CreateLinkList([]int{1, 2})}, want: false},
		{name: "test9", args: args{CreateLinkList([]int{1, 1})}, want: true},
		{name: "e1", args: args{CreateLinkList([]int{1, 1, 0, 0, 1})}, want: false},
		{name: "test10", args: args{CreateLinkList([]int{1, 1, 0, 0, 1, 1})}, want: true},
		{name: "test11", args: args{CreateLinkList([]int{1, 1, 1, 1, 1, 1})}, want: true},
		{name: "test12", args: args{CreateLinkList([]int{1, 1, 1, 1, 0, 1})}, want: false},
		{name: "test13", args: args{CreateLinkList([]int{1, 1, 1, 0, 1, 1, 1})}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
