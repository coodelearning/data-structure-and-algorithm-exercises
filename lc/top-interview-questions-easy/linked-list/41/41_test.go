package _41

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试1",
			args: args{
				node: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  9,
								Next: &ListNode{},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		old := tt.args.node
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			t.Logf("The old node is %+v and the new node is %+v", old, tt.args.node)
		})
	}
}
