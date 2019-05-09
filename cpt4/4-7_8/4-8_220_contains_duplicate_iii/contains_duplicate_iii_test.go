// https://leetcode-cn.com/problems/contains-duplicate-iii/
// 存在重复元素 III
// 给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。
//
// 例 1:
//
// 输入: nums = [1,2,3,1], k = 3, t = 0
// 输出: true
package __8_220_contains_duplicate_iii

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{nums: []int{1, 2, 3, 1}, k: 3, t: 0}, want: true},
		{name: "2", args: args{nums: []int{1, 0, 1, 1}, k: 1, t: 2}, want: true},
		{name: "3", args: args{nums: []int{1, 5, 9, 1, 5, 9}, k: 2, t: 3}, want: false},
		{name: "e1", args: args{nums: []int{1, 0, 1, 1}, k: 0, t: 2}, want: false},
		{name: "e2", args: args{nums: []int{10, 15, 18, 24}, k: 3, t: 3}, want: true},
		{name: "e3", args: args{nums: []int{-5, 5, 5, 5, 5, 15}, k: 6, t: 6}, want: true},
		{name: "u3", args: args{nums: []int{1, 2, 3, 2, 4, 5, 6, 7}, k: 6, t: 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
