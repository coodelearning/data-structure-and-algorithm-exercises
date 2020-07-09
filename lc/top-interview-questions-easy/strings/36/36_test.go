package _6

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "测试3",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "测试4",
			args: args{
				s: "race is car",
			},
			want: false,
		},
		{
			name: "测试5",
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
