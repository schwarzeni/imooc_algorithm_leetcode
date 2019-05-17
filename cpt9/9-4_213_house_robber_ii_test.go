package cpt9

import "testing"

func Test_robii(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]int{1, 2, 3, 1}}, want: 4},
		{name: "test2", args: args{[]int{2, 3, 2}}, want: 3},
		{name: "test3", args: args{[]int{2, 3}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robii(tt.args.nums); got != tt.want {
				t.Errorf("robii() = %v, want %v", got, tt.want)
			}
		})
	}
}
