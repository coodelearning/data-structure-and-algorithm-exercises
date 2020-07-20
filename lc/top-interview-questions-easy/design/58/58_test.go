package _58

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Reset(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.Reset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.Shuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test58(t *testing.T) {
	s := Constructor([]int{1, 2, 3})
	t.Log("--------------0-------------------")
	t.Log(s)
	t.Log("--------------1-------------------")
	t.Log(s.Shuffle())
	t.Log("--------------2-------------------")
	t.Log(s.Reset())
	t.Log("--------------3-------------------")
	t.Log(s.Shuffle())
}
