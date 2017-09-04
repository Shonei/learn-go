package main

import "testing"

func Test_findAnagrams(t *testing.T) {
	type args struct {
		str     string
		anagram string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: args{str: "AbrAcadAbRa", anagram: "cAda"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.str, tt.args.anagram); got != tt.want {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areAnagrams(t *testing.T) {
	type args struct {
		str   string
		check string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test 1", args: args{str: "abc", check: "cba"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAnagrams(tt.args.str, tt.args.check); got != tt.want {
				t.Errorf("areAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
