package main

import "testing"

func Test_atbash(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test 1",
			args: args{str: "abcdefghijklmnopqrstuvwxyzZYXWVUTSRQPONMLKJIHGFEDCBA"},
			want: "zyxwvutsrqponmlkjihgfedcbaABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atbash(tt.args.str); got != tt.want {
				t.Errorf("atbash() = %v, want %v", got, tt.want)
			}
		})
	}
}
