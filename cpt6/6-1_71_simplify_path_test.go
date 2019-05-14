// https://leetcode-cn.com/problems/simplify-path/
// 71. 简化路径
// 请注意，返回的规范路径必须始终以斜杠 / 开头，
// 并且两个目录名之间必须只有一个斜杠 /。
// 最后一个目录名（如果存在）不能以 / 结尾。
// 此外，规范路径必须是表示绝对路径的最短字符串。
package cpt6

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{path: "/a//b////c/d//././/.."}, want: "/a/b/c"},
		{name: "test2", args: args{path: "/a/../../b/../c//.//"}, want: "/c"},
		{name: "test3", args: args{path: "/a/./b/../../c/"}, want: "/c"},
		{name: "test4", args: args{path: "/home//foo/"}, want: "/home/foo"},
		{name: "test5", args: args{path: "/../"}, want: "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
