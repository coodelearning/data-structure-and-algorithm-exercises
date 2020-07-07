package _27

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "2",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			name: "2.1",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "3",
			args: args{
				digits: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1},
		},
		{
			name: "4",
			args: args{
				digits: []int{2, 4, 9, 3, 9},
			},
			want: []int{2, 4, 9, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
