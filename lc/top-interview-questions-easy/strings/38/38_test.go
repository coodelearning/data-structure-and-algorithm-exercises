package _38

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "测试2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "测试3",
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			name: "测试4",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr(%v,%v) = %v, want %v", tt.args.haystack, tt.args.needle, got, tt.want)
			}
		})
	}
}
