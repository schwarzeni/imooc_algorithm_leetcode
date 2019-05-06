package minimum_size_subarray_sum

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	//type args struct {
	//	s    int
	//	nums []int
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
	//		if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
	//			t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	//
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(1000, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(1000, []int{}))
}
