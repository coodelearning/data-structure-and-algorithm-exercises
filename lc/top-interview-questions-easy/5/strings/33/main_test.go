package _33

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "测试2",
			args: args{
				x: -321,
			},
			want: -123,
		},
		{
			name: "测试3",
			args: args{
				x: 12000,
			},
			want: 21,
		},
		{
			name: "测试4",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "测试5",
			args: args{
				x: 901000,
			},
			want: 109,
		},
		{
			name: "测试6",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("the int is %v reverse() = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
