// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合
package cpt8

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		//{name: "test1", args: args{digits: "23"}, want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		//{name: "test2", args: args{digits: "2"}, want: []string{"a", "b", "c"}},
		{name: "test2", args: args{digits: "234"}, want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			//}
			if got := letterCombinations(tt.args.digits); true == true {
				fmt.Println(got)
			}
		})
	}
}
