package fib

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int64
	}
	var tests = []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{n: 1}, want: 1},
		{name: "2", args: args{n: 2}, want: 1},
		{name: "3", args: args{n: 3}, want: 2},
		{name: "4", args: args{n: 4}, want: 3},
		{name: "5", args: args{n: 5}, want: 5},
		{name: "6", args: args{n: 6}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
