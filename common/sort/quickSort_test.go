package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{3, 44, 38, 5, 47, 15, 36, 26, 26, 27, 2, 46, 4, 19, 50, 48}
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{a, 0, len(a) - 1}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
