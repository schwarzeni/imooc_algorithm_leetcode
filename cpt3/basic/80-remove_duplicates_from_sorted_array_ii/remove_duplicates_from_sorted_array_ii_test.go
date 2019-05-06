package _80_remove_duplicates_from_sorted_array_ii

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{name: "", args: args{nums: []int{}}, expected: 0},
		{name: "", args: args{nums: []int{1, 1, 2, 2, 3, 3}}, expected: 6},
		{name: "", args: args{nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3}}, expected: 6},
		{name: "", args: args{nums: []int{1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3}}, expected: 6},
		{name: "", args: args{nums: []int{1, 1, 1, 2, 2, 3}}, expected: 5},
		{name: "", args: args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, expected: 7},
		{name: "", args: args{nums: []int{0, 0, 0, 0, 0, 0}}, expected: 2},
		{name: "", args: args{nums: []int{0, 1, 2, 3, 4, 5, 6}}, expected: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print(tt.args.nums, " ")
			result := removeDuplicates(tt.args.nums)
			fmt.Println(tt.args.nums)
			if result != tt.expected {
				t.Errorf("expect %d, but get %d", tt.expected, result)
			}
		})
	}
}
