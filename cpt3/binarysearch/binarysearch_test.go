// 二分查找法
package binarysearch

import "testing"

func Test_binarySearch(t *testing.T) {
	testArr := []int{0,1,2,3,4,5,6,7,8,9,10}
	type args struct {
		arr    []int
		size   int
		target int
	}
	type test struct {
		name    string
		args    args
		wantIdx int
	}
	var tests []test
	for idx, data := range testArr {
		tests = append(tests, test{name:"", args:args{arr:testArr, size:len(testArr), target:data}, wantIdx:idx} )
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIdx := binarySearch(tt.args.arr, tt.args.size, tt.args.target); gotIdx != tt.wantIdx {
				t.Errorf("binarySearch() = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
