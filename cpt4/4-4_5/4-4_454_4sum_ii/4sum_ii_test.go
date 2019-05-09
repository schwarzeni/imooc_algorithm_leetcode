package __4_454_4sum_ii

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{A: []int{1, 2}, B: []int{-2, -1}, C: []int{-1, 2}, D: []int{0, 2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
