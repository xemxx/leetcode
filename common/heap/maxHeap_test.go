package heap

import "testing"

func Test_maxHeap_Push(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &maxHeap{
				data: tt.fields.data,
			}
			c.Push(tt.args.x)
		})
	}
}
