package __7_217_contains_duplicate

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{nums: []int{1, 2, 3, 1}}, want: true},
		{name: "2", args: args{nums: []int{1, 0, 1, 1}}, want: true},
		{name: "3", args: args{nums: []int{1, 2, 3}}, want: false},
		{name: "3", args: args{nums: []int{1, 1, 2, 3}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
