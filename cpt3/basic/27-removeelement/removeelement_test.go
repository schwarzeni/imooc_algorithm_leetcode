package _27_removeelement

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{nums: []int{0, 1, 0, 3, 12}, val: 0}},
		{name: "", args: args{nums: []int{0, 0, 0, 0}, val: 0}},
		{name: "", args: args{nums: []int{0, 1, 3, 0}, val: 0}},
		{name: "", args: args{nums: []int{1, 2, 3, 0}, val: 0}},
		{name: "", args: args{nums: []int{1, 2, 3, 0, 0}, val: 0}},
		{name: "", args: args{nums: []int{1, 0, 1}, val: 0}},
		{name: "", args: args{nums: []int{0}, val: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print(tt.args.nums, " ")
			fmt.Print(removeElement(tt.args.nums, tt.args.val), " ")
			fmt.Println(tt.args.nums)
		})
	}
}
