package _48

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: true,
		},
		{
			name: "测试2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "测试3",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:   -1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: true,
		},
		{
			name: "测试4",
			args: args{
				root: &TreeNode{},
			},
			want: true,
		},
		{
			name: "测试5",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   20,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "测试6",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		//if tt.name != "测试6" {
		//	continue
		//}
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
