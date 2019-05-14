package cpt6

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		//{name: "test1", args: args{[]int{1, NULL_NODE, 2, 3}}},
		{name: "test1", args: args{[]int{15, 4, 7, 3, NULL_NODE, 2, NULL_NODE, -1, NULL_NODE, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTree(tt.args.arr); true == true {
				fmt.Println(got)
			}
		})
	}
}
