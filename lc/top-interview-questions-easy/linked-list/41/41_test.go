package linked_list

import (
	"data-structure-and-algorithm-exercises/lc/top-interview-questions-easy/linked-list/base"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *base.ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试1",
			args: args{
				node: &base.ListNode{
					Val: 4,
					Next: &base.ListNode{
						Val: 5,
						Next: &base.ListNode{
							Val: 1,
							Next: &base.ListNode{
								Val:  9,
								Next: &base.ListNode{},
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
