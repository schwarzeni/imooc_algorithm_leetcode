// https://leetcode-cn.com/problems/valid-anagram/
// 有效的字母异位词
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 输入: s = "rat", t = "car"
// 输出: false
package valid_anagram

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "anagram", t: "nagaram"}, want: true},
		{name: "test2", args: args{s: "rat", t: "car"}, want: false},
		{name: "test3", args: args{s: "ratttt", t: "car"}, want: false},
		{name: "test4", args: args{s: "", t: ""}, want: true},
		{name: "test4", args: args{s: "car", t: "ratttt"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
