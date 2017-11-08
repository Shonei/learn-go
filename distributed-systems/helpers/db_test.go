package hlp

import (
	"testing"
)

func Test_getCreateQuery(t *testing.T) {
	type args struct {
		table Table
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1",
			args{Table{"Shonei", []Pair{Pair{"hi", ""}, Pair{"world", ""}}}},
			"create table Shonei(hi text,world text);"},
		{"Test 2",
			args{Table{"Shyl", []Pair{Pair{"hi", ""}, Pair{"world", ""}}}},
			"create table Shyl(hi text,world text);"},
		{"Test 3",
			args{Table{"Kipendy", []Pair{Pair{"hi", ""}, Pair{"world", ""}}}},
			"create table Shyl(hi text,world text);"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCreateQuery(tt.args.table); got != tt.want {
				t.Errorf("getCreateQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInsertQuery(t *testing.T) {
	type args struct {
		keys []Pair
		user string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInsertQuery(tt.args.keys, tt.args.user); got != tt.want {
				t.Errorf("getInsertQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyCheck(t *testing.T) {
	type args struct {
		stored []Pair
		user   []Pair
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1",
			args{[]Pair{Pair{"a", "sdh"}, Pair{"b", "sgf"}, Pair{"c", "dagsdg"}},
				[]Pair{Pair{"a", "sdfh"}, Pair{"b", "fgh"}, Pair{"c", "sghsgg"}}},
			true},
		{"test1",
			args{[]Pair{Pair{"a", "sdh"}, Pair{"b", "sgf"}, Pair{"c", "dagsdg"}},
				[]Pair{Pair{"x", "sdfh"}, Pair{"b", "fgh"}, Pair{"c", "sghsgg"}}},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyCheck(tt.args.stored, tt.args.user); got != tt.want {
				t.Errorf("keyCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
