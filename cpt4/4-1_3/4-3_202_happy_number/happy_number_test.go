package happy_number

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{n: 19}, want: true},
		{name: "test1", args: args{n: 1}, want: true},
		{name: "test1", args: args{n: 100}, want: true},
		{name: "test1", args: args{n: 56}, want: false},
		{name: "test1", args: args{n: 15}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
