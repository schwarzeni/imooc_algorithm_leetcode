// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号
// 输入: "([)]"
// 输出: false
package cpt6

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "()"}, want: true},
		{name: "test2", args: args{s: "("}, want: false},
		{name: "test3", args: args{s: "(])"}, want: false},
		{name: "test4", args: args{s: "([])"}, want: true},
		{name: "test5", args: args{s: "([{])"}, want: false},
		{name: "test6", args: args{s: "{([{}])}"}, want: true},
		{name: "test7", args: args{s: "{([{}])"}, want: false},
		{name: "test8", args: args{s: ""}, want: true},
		{name: "e1", args: args{s: "()[]{}"}, want: true},
		{name: "e2", args: args{s: "{[]}"}, want: true},
		{name: "e3", args: args{s: "([)]"}, want: false},
		{name: "e4", args: args{s: "{}[{{}][]("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
