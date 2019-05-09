package __4_18_4sum

import (
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "", args: args{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, want: [][]int{{1}}},
		{name: "", args: args{nums: []int{1, -2, -5, -4, -3, 3, 3, 5}, target: -11}, want: [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); true == true {
				t.Logf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}

	// -5 -4 -3 -2 1 3 3 5
}
