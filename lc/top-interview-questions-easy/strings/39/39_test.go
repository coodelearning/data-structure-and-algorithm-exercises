package _39

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "测试2",
			args: args{
				n: 4,
			},
			want: "1211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				// if got := countAndSayBetter(tt.args.n); got != tt.want {
				t.Errorf("countAndSay(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
