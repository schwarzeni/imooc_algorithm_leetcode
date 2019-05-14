// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
package cpt6

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{tokens: []string{"2", "1", "+", "3", "*"}}, want: 9},
		{name: "test2", args: args{tokens: []string{"4", "13", "5", "/", "+"}}, want: 6},
		{name: "test3", args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
