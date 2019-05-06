// 数组中仅有 0，1，2 三个元素，请对其进行排序
// https://leetcode-cn.com/problems/sort-colors/submissions/
package _80_sort_colors

import (
	"fmt"
	"testing"
)

func Test_sortColorsBetter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColorsBetter(tt.args.nums)
		})
	}
	nums := []int{1, 2, 0}
	sortColorsBetter(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 1, 2, 0, 1, 2}
	sortColorsBetter(nums)
	fmt.Println(nums)

	nums = []int{0, 0, 0}
	sortColorsBetter(nums)
	fmt.Println(nums)

	nums = []int{2, 0, 2, 1, 1, 0}
	sortColorsBetter(nums)
	fmt.Println(nums)

	nums = []int{0}
	sortColorsBetter(nums)
	fmt.Println(nums)

	nums = []int{2, 0}
	sortColorsBetter(nums)
	fmt.Println(nums)
}
