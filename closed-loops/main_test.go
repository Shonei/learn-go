package main

import "testing"

func Test_getLoopCount(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Hello", args: args{i: 57}, want: 1},
		{name: "Hello", args: args{i: 54}, want: 1},
		{name: "Hello", args: args{i: 48}, want: 1},
		{name: "Hello", args: args{i: 56}, want: 2},
		{name: "Hello", args: args{i: 5}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLoopCount(tt.args.i); got != tt.want {
				t.Errorf("getLoopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countLoops(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Heloo2", args: args{number: 666}, want: 3},
		{name: "Heloo2", args: args{number: 685}, want: 3},
		{name: "Heloo2", args: args{number: 686}, want: 4},
		{name: "Heloo2", args: args{number: 111}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLoops(tt.args.number); got != tt.want {
				t.Errorf("countLoops() = %v, want %v", got, tt.want)
			}
		})
	}
}
