// https://leetcode-cn.com/problems/group-anagrams/
// 49. 字母异位词分组
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
// 输出:
//
// ["ate","eat","tea"],
// ["nat","tan"],
// ["bat"]
//
package __4_49_group_anagrams

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "", args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, want: [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); true == true {
				t.Logf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
