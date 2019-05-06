package minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{name: "边界1", args: args{s: "ABCDDDDCDBA", t: "ABC"}, want: "ABC"},
		{name: "边界2", args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
		{name: "边界3", args: args{s: "AAAAAAAAAAAA", t: "AA"}, want: "AA"},
		{name: "边界4", args: args{s: "ABBBBBBBBBBA", t: "AA"}, want: "ABBBBBBBBBBA"},
		{name: "边界5", args: args{s: "ABBBBBcBBBBA", t: "AcA"}, want: "ABBBBBcBBBBA"},
		{name: "边界6", args: args{s: "ABBBBBcBBBBA", t: "Ad"}, want: ""},
		{name: "边界7", args: args{s: "AAAAAAAAAAAAA", t: "B"}, want: ""},
		{name: "边界8", args: args{s: "bba", t: "ab"}, want: "ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
