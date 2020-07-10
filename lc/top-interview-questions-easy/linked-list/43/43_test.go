package linked_list

import (
	"data-structure-and-algorithm-exercises/lc/top-interview-questions-easy/linked-list/base"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *base.ListNode
	}
	tests := []struct {
		name string
		args args
		want *base.ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
