package kth_largest_element_in_an_array

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	//type args struct {
	//	nums []int
	//	k    int
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want int
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
	//			t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	testArr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargestII(testArr, 2))

	testArr = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargestII(testArr, 4))
}
