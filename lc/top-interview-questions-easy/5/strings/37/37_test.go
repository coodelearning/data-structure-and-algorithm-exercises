package _37

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "测试2",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "测试3",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "测试4",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "测试5",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi(%v) = %v, want %v", tt.args.str, got, tt.want)
			}
		})
	}
}
