package __6_477_number_of_boomerangs

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{points: [][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}}}, want: 2},
		{name: "", args: args{points: [][]int{[]int{0, 0}, []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
