package _43

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "测试1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: &ListNode{},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val:  0,
				Next: &ListNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//got := reverseListIteration(tt.args.head)
			got := reverseListRecursive(tt.args.head)
			for got.Next != nil {
				t.Log("The value is ", got.Val)
				got = got.Next
			}
			//if got := reverseListRecursive(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			//t.Errorf("reverseList() = %v, want %v", got, tt.want)
		})
	}
}
