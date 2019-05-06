// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 仅含小写英文字母
package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{name: "test1", args: args{s: "cbaebabacd", p: "abc"}, want: []int{0, 6}},
		//{name: "test2", args: args{s: "abab", p: "ab"}, want: []int{0, 1, 2}},
		//{name: "边界1", args: args{s: "aaaa", p: "aaaaaa"}, want: []int{}},
		//{name: "边界2", args: args{s: "aaaa", p: "aaaa"}, want: []int{0}},
		{name: "边界2", args: args{s: func() string {
			str := ""
			for i := 0; i < 1500; i++ {
				str += "a"
			}
			return str
		}(), p: func() string {
			str := ""
			for i := 0; i < 1500; i++ {
				str += "a"
			}
			return str
		}()}, want: []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
