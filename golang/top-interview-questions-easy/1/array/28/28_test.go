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
			t.Log("test nums", tt.name, tt.args.nums)
			moveZeroes(tt.args.nums)
			betterMoveZeroes(tt.args.nums)
			t.Log("moveZeroes", tt.name, tt.args.nums)
			t.Log("betterMoveZeroes", tt.name, tt.args.nums)
		})
	}
}
