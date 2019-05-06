package _26_remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{nums: []int{}}},
		{name: "", args: args{nums: []int{1, 2, 3, 4, 5}}},
		{name: "", args: args{nums: []int{0, 0, 1}}},
		{name: "", args: args{nums: []int{0, 1, 1}}},
		{name: "", args: args{nums: []int{0, 0, 0, 0}}},
		{name: "", args: args{nums: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}}},
		{name: "", args: args{nums: []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print(tt.args.nums, " ")
			fmt.Print(removeDuplicates(tt.args.nums), " ")
			fmt.Println(tt.args.nums)
		})
	}
}
