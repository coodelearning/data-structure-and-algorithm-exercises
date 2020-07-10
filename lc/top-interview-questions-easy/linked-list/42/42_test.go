package linked_list

import (
	"data-structure-and-algorithm-exercises/lc/top-interview-questions-easy/linked-list/base"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *base.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *base.ListNode
	}{
		{
			name: "测试1",
			args: args{
				head: &base.ListNode{
					Val: 1,
					Next: &base.ListNode{
						Val: 2,
						Next: &base.ListNode{
							Val: 3,
							Next: &base.ListNode{
								Val: 4,
								Next: &base.ListNode{
									Val:  5,
									Next: &base.ListNode{},
								},
							},
						},
					},
				},
				n: 2,
			},
			want: &base.ListNode{
				Val: 1,
				Next: &base.ListNode{
					Val: 2,
					Next: &base.ListNode{
						Val: 3,
						Next: &base.ListNode{
							Val:  5,
							Next: &base.ListNode{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
