package __4_1_two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "test2", args: args{nums: []int{2, 7, 11, 15}, target: 10}, want: []int{-1, -1}},
		{name: "test3", args: args{nums: []int{7, 11, 2, 15}, target: 9}, want: []int{0, 2}},
		{name: "test4", args: args{nums: []int{1, 10, 100, 10000, 1000000}, target: 10100}, want: []int{2, 3}},
		{name: "test5", args: args{nums: []int{}, target: 9}, want: []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Logf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
