package cpt8

import (
	"fmt"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{matrix: [][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}}},
		{name: "test2", args: args{matrix: [][]int{
			{5, 5, 5, 5, 1},
			{5, 2, 5, 2, 1},
			{1, 5, 1, 2, 1},
			{1, 2, 2, 2, 1},
			{1, 1, 1, 1, 1},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pacificAtlantic(tt.args.matrix); true == true {
				fmt.Println(got, len(got))
			}
		})
	}
}
