package cpt5

import (
	"testing"
)

func TestCreateLinkList(t *testing.T) {
	type args struct {
		dataArr []int
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *ListNode
	}{
		{name: "test1", args: args{[]int{1, 2, 3, 4, 5, 6}}},
		{name: "test2", args: args{[]int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHeader := CreateLinkList(tt.args.dataArr); true == true {
				PrintLinkList(gotHeader)
			}
		})
	}
}
