// https://leetcode-cn.com/problems/contains-duplicate-ii/
// 存在重复元素 II
// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
// 示例 1:
// 入: nums = [1,2,3,1], k = 3
// 出: true
package __7_219_contains_duplicate_ii

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{nums: []int{1, 2, 3, 1}, k: 3}, want: true},
		{name: "2", args: args{nums: []int{1, 0, 1, 1}, k: 1}, want: true},
		{name: "3", args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2}, want: false},
		{name: "4", args: args{nums: []int{1}, k: 2}, want: false},
		{name: "5", args: args{nums: []int{1, 1}, k: 1}, want: true},
		{name: "6", args: args{nums: []int{1, 1}, k: 0}, want: false},
		{name: "7", args: args{nums: []int{1, 2}, k: 1}, want: false},
		{name: "8", args: args{nums: []int{1, 2}, k: 2}, want: false},
		{name: "9", args: args{nums: []int{99, 99}, k: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
