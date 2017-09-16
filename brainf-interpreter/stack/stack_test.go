package stack

import "testing"

func TestStack_Push(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		s    *Stack
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.v)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Pop()
		})
	}
}
