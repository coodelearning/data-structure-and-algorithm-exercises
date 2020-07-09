package _40

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "测试2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "测试3",
			args: args{
				strs: []string{"a"},
			},
			want: "a",
		},
		{
			name: "测试4",
			args: args{
				strs: []string{"c", "c"},
			},
			want: "c",
		},
		{
			name: "测试5",
			args: args{
				strs: []string{"aa", "aa"},
			},
			want: "aa",
		},
		{
			name: "测试6",
			args: args{
				strs: []string{"aa", "a"},
			},
			want: "a",
		},
		{
			name: "测试7",
			args: args{
				strs: []string{"a", "a", "b"},
			},
			want: "",
		},
		{
			name: "测试8",
			args: args{
				strs: []string{"abab", "aba", "abc"},
			},
			want: "ab",
		},
		{
			name: "测试9",
			args: args{
				strs: []string{"aaa", "aa", "aaa"},
			},
			want: "aa",
		},
	}
	for _, tt := range tests {
		//if tt.name != "测试9" {
		//	continue
		//}
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %v, want %v", tt.args.strs, got, tt.want)
			}
		})
	}
}
