// https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 根据字符出现频率排序
package sort_characters_by_frequency

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{s: "tree"}, want: "eert"},
		{name: "test2", args: args{s: "Aabb"}, want: "bbAa"},
		{name: "test3", args: args{s: "aaabbbccc"}, want: "aaabbbccc"},
		{name: "test4", args: args{s: (func() string {
			total := 100000000
			str := make([]int32, total)
			for i := 0; i < total; i++ {
				if i%2 == 0 {
					str[i] = 'a'
				} else {
					str[i] = 'b'
				}
				if i == total-1 {
					str[i] = 'a'
				}
			}
			return string(str)
		})()}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := frequencySort(tt.args.s); got != tt.want {
			//	t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			//}
			got := frequencySort(tt.args.s)
			t.Log(got)
		})
	}
}
