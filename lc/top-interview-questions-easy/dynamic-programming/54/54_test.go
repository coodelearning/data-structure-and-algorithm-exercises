package _54

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				n: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
