// https://leetcode-cn.com/problems/max-points-on-a-line/
// 直线上最多的点数
package __6_149_max_points_on_a_line

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int // 参数结构
	}
	tests := []struct {
		name string // 测试名称
		args args   // 传入参数
		want int    // 预期数据
	}{
		{name: "没有点", args: args{points: [][]int{}}, want: 0},
		{name: "只有一个点", args: args{points: [][]int{
			[]int{1, 1},
		}}, want: 1},
		{name: "只有两个点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
		}}, want: 2},
		{name: "有两个(1,1)的点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, 1},
		}}, want: 3},
		{name: "有两个点存在重复的情况(1,1), (0,0)", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, 1},
			[]int{0, 0},
			[]int{2, 3},
		}}, want: 4},
		{name: "正常测试1", args: args{points: [][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{3, 3},
		}}, want: 3},
		{name: "正常测试2", args: args{points: [][]int{
			[]int{1, 1},
			[]int{3, 2},
			[]int{5, 3},
			[]int{4, 1},
			[]int{2, 3},
			[]int{1, 4},
		}}, want: 4},
		{name: "三个点相同的点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{1, 1},
			[]int{1, 1},
		}}, want: 3},
		{name: "斜率为0", args: args{points: [][]int{
			[]int{2, 3},
			[]int{3, 3},
			[]int{-5, 3},
			[]int{2, 1},
		}}, want: 3},
		{name: "斜率为无穷大", args: args{points: [][]int{
			[]int{3, -1},
			[]int{3, 2},
			[]int{3, 1},
			[]int{2, 1},
		}}, want: 3},
		{name: "k除不尽", args: args{points: [][]int{
			[]int{84, 250},
			[]int{0, 0},
			[]int{1, 0},
			[]int{0, -70},
			[]int{0, -70},
			[]int{1, -1},
			[]int{21, 10},
			[]int{42, 90},
			[]int{-42, -230},
		}}, want: 6},
		{name: "浮点数精度丢失", args: args{points: [][]int{
			[]int{94911152, 94911151},
			[]int{0, 0},
			[]int{94911151, 94911150},
		}}, want: 2},
	}
	for _, tt := range tests {
		//if tt.name == "10" {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
		//}

	}
}
