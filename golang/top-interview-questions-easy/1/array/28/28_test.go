package _28

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			t.Log(tt.name, tt.args.nums)
		})
	}
}
