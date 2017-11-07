package hlp

import (
	"testing"
)

// func TestCreateTable(t *testing.T) {
// 	type args struct {
// 		j     []byte
// 		db    *sql.DB
// 		uname string
// 		email string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want Responce
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := CreateTable(tt.args.j, tt.args.db, tt.args.uname, tt.args.email); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CreateTable() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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

// func Test_errHandler(t *testing.T) {
// 	type args struct {
// 		err error
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			errHandler(tt.args.err)
// 		})
// 	}
// }
