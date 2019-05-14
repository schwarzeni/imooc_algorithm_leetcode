// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 347. 前K个高频元素
package cpt6

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{name: "test2", args: args{nums: []int{1}, k: 1}, want: []int{1}},
		{name: "test3", args: args{nums: []int{4, 4, 5, 5, 5, 4, 3, 4, 5, 5, 3, 6}, k: 3}, want: []int{5, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
