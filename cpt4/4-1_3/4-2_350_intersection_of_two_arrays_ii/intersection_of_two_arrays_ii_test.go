// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// 两个数组的交集 II
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
package intersection_of_two_arrays_ii

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2, 2}},
		{name: "test2", args: args{nums1: []int{1, 2, 2, 3}, nums2: []int{4, 5}}, want: []int{}},
		{name: "test3", args: args{nums1: []int{}, nums2: []int{}}, want: []int{}},
		{name: "test4", args: args{nums1: []int{1, 2}, nums2: []int{}}, want: []int{}},
		{name: "test5", args: args{nums1: []int{}, nums2: []int{1, 2}}, want: []int{}},
		{name: "test6", args: args{nums1: []int{2, 2, 3}, nums2: []int{1, 2, 2, 1, 3}}, want: []int{2, 2, 3}},
		{name: "test7", args: args{nums1: []int{4, 9, 5, 9, 4}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{9, 4, 9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, tt.want
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(sortArr(got, tt.want)) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortArr(arr1 []int, arr2 []int) ([]int, []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	return arr1, arr2
}
