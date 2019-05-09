// https://leetcode-cn.com/problems/word-pattern/
// 单词模式
package word_pattern

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: args{pattern: "abba", str: "dog cat cat dog"}, want: true},
		{name: "test", args: args{pattern: "abba", str: "dog cat cat fish"}, want: false},
		{name: "test", args: args{pattern: "aaaa", str: "dog cat cat dog"}, want: false},
		{name: "test", args: args{pattern: "abba", str: "dog dog dog dog"}, want: false},
		{name: "test", args: args{pattern: "aaaa", str: "dog dog dog dog"}, want: true},
		{name: "test", args: args{pattern: "", str: "dog dog dog dog"}, want: false},
		{name: "test", args: args{pattern: "aaaa", str: ""}, want: false},
		{name: "test", args: args{pattern: "abcaccd", str: "dog cat bird dog bird bird tiger"}, want: true},
		{name: "test", args: args{pattern: "a", str: "abb"}, want: true},
		{name: "test", args: args{pattern: "a", str: "abb vc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
