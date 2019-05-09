// https://leetcode-cn.com/problems/intersection-of-two-arrays/
// 求两个数组的交集
package intersection_of_two_arrays

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2}},
		{name: "test2", args: args{nums1: []int{1, 2, 2, 3}, nums2: []int{4, 5}}, want: []int{}},
		{name: "test3", args: args{nums1: []int{}, nums2: []int{}}, want: []int{}},
		{name: "test4", args: args{nums1: []int{1, 2}, nums2: []int{}}, want: []int{}},
		{name: "test5", args: args{nums1: []int{}, nums2: []int{1, 2}}, want: []int{}},
		{name: "test6", args: args{nums1: []int{2, 2, 3}, nums2: []int{1, 2, 2, 1, 3}}, want: []int{2, 3}},
		{name: "test7", args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
