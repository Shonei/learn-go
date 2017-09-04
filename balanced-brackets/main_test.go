package main

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test 1",
			args: args{str: "()[]{}(([])){[()][]}"},
			want: true},
		{name: "Test 2",
			args: args{str: "())[]{}"},
			want: false},
		{name: "Test 2",
			args: args{str: "[(])"},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.str); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
