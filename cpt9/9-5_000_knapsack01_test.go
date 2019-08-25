// 0-1背包问题
// 填表
package cpt9

import "testing"

func Test_knapsack01(t *testing.T) {
	type args struct {
		v []int
		w []int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "背包问题1", args: args{v: []int{6, 10, 12}, w: []int{1, 2, 3}, c: 5}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack01(tt.args.v, tt.args.w, tt.args.c); got != tt.want {
				t.Errorf("knapsack01() = %v, want %v", got, tt.want)
			}
		})
	}
}
