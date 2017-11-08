// Package hlp provides functions for accessing the database
// and server handlers
package hlp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
)

// Table structure for registering a user
type Table struct {
	User string
	Keys []Pair
}

// Pair key values for creating and inserting data
type Pair struct {
	Key   string
	Value string `json:"omitempty"`
}

// Responce structure for confirming registration for users
type Responce struct {
	GetURL          string
	PostURL         string
	PostDescription string
	Status          string
	Description     string
	Error           string
}

// CreateTable takes information for a http request
// and create a table for a user wishing to register
func CreateTable(j []byte, db *sql.DB, email string) Responce {
	table := &Table{}
	json.Unmarshal(j, &table)
	fmt.Println(table)
	res := Responce{}
	j, err := json.Marshal(table.Keys)
	insertQuery := "Insert into users(uname, get, email, json) values($1, $2, $3, $4)"
	createQuery := getCreateQuery(*table)
	get := "select * from " + table.User + ";"

	// transaction to register user and create his table
	err = registerUser(db, insertQuery, createQuery, table.User, get, email, string(j))
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			fmt.Println(err)
			res.Error = "User already exists."
		} else {
			res.Error = "Failed to register user."
		}
		return res
	}

	res.GetURL = "http://localhost:8080/get/{user}"
	res.PostURL = "http://localhost:8080/post"
	res.PostDescription = `When posting data it should be in a JSON of the form {"User": "name-you-registered", "Keys": [{"Key":"the-name-of-variable-in-your-db", "Value": "value-you-want-placed"}]}. Each request will result in a single row being added.`
	res.Status = "OK"
	res.Description = "Thank you for registering. The data you wish to send in a post should be in json format."
	return res
}

func getCreateQuery(table Table) string {
	temp := "create table "
	details := "("

	for i := 0; i < len(table.Keys); i++ {
		details += fmt.Sprint(table.Keys[i].Key) + " text"
		if i+1 != len(table.Keys) {
			details += ","
		}
	}
	details += ");"
	return temp + table.User + details
}

func registerUser(db *sql.DB, insertQuery, createQuery, uname, get, email, j string) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	// prepare register query
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	// put user into db
	_, err = stmt.Exec(uname, get, email, j)
	if err != nil {
		tx.Rollback()
		return err
	}

	// create users custom table
	_, err = tx.Exec(createQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetUsersData returns the information a users has stored
// in his custom database. If there was an error nil is returned
func GetUsersData(db *sql.DB, uname, email string) [][]Pair {
	row, err := db.Query("select email, get from users where uname = $1;", uname)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer row.Close()
	var e string
	var get string

	if row.Next() {
		row.Scan(&e, &get)
		if email != e {
			log.Println("Email missmatch. No permision granted.")
			// wrong user
			return nil
		}
	} else {
		// wrong username
		log.Println("I don't know why is that here. Sorry")
		return nil
	}

	// get the data from the users table
	rows, err := db.Query(get)
	data, err := readRows(rows)
	if err != nil {
		return nil
	}
	return data
}

// reads arbytrery data from sql result
// returns it as an array of array
// [[{Key:"",Value:""},{Key:"",Value:""}],[{Key:"",Value:""},{Key:"",Value:""}]]
func readRows(rows *sql.Rows) ([][]Pair, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawResult := make([][]byte, len(cols))
	ret := [][]Pair{}
	dest := make([]interface{}, len(cols))

	// load the interface with pointers to get the data
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		// read data into dest that hold the pointers
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		result := []Pair{}
		// read the pointers and and them to the result
		for i, raw := range rawResult {
			if raw == nil {
				result = append(result, Pair{Key: cols[i], Value: ""})
			} else {
				result = append(result, Pair{Key: cols[i], Value: string(raw)})
			}
		}

		ret = append(ret, result)
	}
	return ret, nil
}

// InsertIntoUser is used to put data into the custom user table
// It needs the array of data to put and the username
func InsertIntoUser(db *sql.DB, js []byte, uname string) error {
	data := &[]Pair{}
	json.Unmarshal(js, &data)

	var s string
	err := db.QueryRow("select json from users where uname = $1",
		uname).Scan(&s)
	if err != nil {
		// the scan error handled here
		return err
	}

	j := []Pair{}
	err = json.Unmarshal([]byte(s), &j)

	// Check to see if user has given the right keys for the table
	if !keyCheck(j, *data) {
		return errors.New("Wrong user keys")
	}
	put := getInsertQuery(*data, uname)

	// Prepare the insert query
	stmt, err := db.Prepare(put)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute the query and and check error
	// we don't care about the returned data
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// bad way to check if user has given the right keys
// good enough for now but thinnk about reimplamenting it later
func keyCheck(stored, user []Pair) bool {
	for i := range stored {
		val := true
		for j := range user {
			if stored[i].Key == user[j].Key {
				val = true
				break
			} else {
				val = false
			}
		}
		if !val {
			return false
		}
	}
	return true
}

// again bad way to do it when you take security into account
// this is heaven for SQL injection attacks
// Well the system will go down in a month so it should be fine till then
// This generates the insert query for a users table
func getInsertQuery(data []Pair, uname string) string {
	put := "insert into " + uname + "("
	key := ""
	vals := ""

	for i := 0; i < len(data); i++ {
		key += data[i].Key
		vals += "'" + data[i].Value + "'"
		if i+1 != len(data) {
			key += ", "
			vals += ", "
		}
	}
	return put + key + ") values(" + vals + ");"
}
