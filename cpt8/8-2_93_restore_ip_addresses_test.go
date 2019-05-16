// https://leetcode-cn.com/problems/restore-ip-addresses/
// 93. 复原IP地址
// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]
package cpt8

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{s: "25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "test2",
			args: args{s: "25525525522"},
			want: []string{"255.255.255.22"},
		},
		{
			name: "test3",
			args: args{s: "255255111352"},
			want: []string{},
		},
		{
			name: "test4",
			args: args{s: "2552"},
			want: []string{"2.5.5.2"},
		},
		{
			name: "test5",
			args: args{s: "255"},
			want: []string{},
		},
		{
			name: "e1",
			args: args{s: "172162541"},
			want: []string{"17.216.25.41", "17.216.254.1", "172.16.25.41", "172.16.254.1", "172.162.5.41", "172.162.54.1"},
		},
		{
			name: "e2",
			args: args{s: "000256"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
