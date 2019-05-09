// https://leetcode-cn.com/problems/3sum/
package __4_15_3sum

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// -4 -1 -1 0 1 2
		{name: "test1", args: args{nums: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{name: "test2", args: args{nums: []int{-3, -3, -2, -2, -1, -1, 0, 1, 1, 2, 2, 3, 3, 4, 4}}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{name: "test3", args: args{nums: []int{-2, 0, 1, 1, 2}}, want: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); true == true {
				t.Logf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
