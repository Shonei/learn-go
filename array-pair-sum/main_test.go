package main

import (
	"reflect"
	"testing"
)

func Test_getSumPairs(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want []pair
	}{
		{name: "Test 1",
			args: args{arr: []int{3, 4, 5, 4, 4}, sum: 8},
			want: []pair{{3, 5}, {4, 4}, {4, 4}, {4, 4}}},
		{name: "Test 2",
			args: args{arr: []int{3, 4, 5, 6, 7}, sum: 10},
			want: []pair{{3, 7}, {4, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumPairs(tt.args.arr, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
