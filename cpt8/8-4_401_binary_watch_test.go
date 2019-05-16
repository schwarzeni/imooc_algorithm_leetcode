// https://leetcode-cn.com/problems/binary-watch/
// 401. 二进制手表
// 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
// 每个 LED 代表一个 0 或 1，最低位在右侧。
// 输入: n = 1
// 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
package cpt8

import (
	"fmt"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{num: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.num); true == true {
				fmt.Println(got)
			}
		})
	}
}
