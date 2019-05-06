package pow

import (
	"testing"
)

func Test_pow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{name: "1", args: args{x: 4.0, n: 3}, wantResult: 64.0},
		{name: "2", args: args{x: 10.0, n: 7}, wantResult: 10000000.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := betterPow(tt.args.x, tt.args.n); gotResult != tt.wantResult {
				t.Errorf("betterPow() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_normalPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{name: "1", args: args{x: 4.0, n: 3}, wantResult: 64.0},
		{name: "2", args: args{x: 10.0, n: 7}, wantResult: 10000000.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := normalPow(tt.args.x, tt.args.n); gotResult != tt.wantResult {
				t.Errorf("normalPow() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func BenchmarkTest_betterPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		betterPow(2233.0, 3322)
	}
}

func BenchmarkTest_normalPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalPow(2233.0, 3322)
	}
}

