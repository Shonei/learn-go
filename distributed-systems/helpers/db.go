package hlp

import (
	"database/sql"
	"encoding/json"
	"fmt"

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
	Value string
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
	res := Responce{}
	insertQuery := "Insert into users(uname, get, insertQ, email) values($1, $2, $3, $4)"
	createQuery := getCreateQuery(*table)
	get := "select * from " + table.User + ";"
	insert := getInsertQuery(table.Keys, table.User)

	// transaction to register user and create his table
	err := registerUser(db, insertQuery, createQuery, table.User, get, insert, email)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
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

func getInsertQuery(keys []Pair, user string) string {
	v := ""
	k := ""
	for i := 0; i < len(keys); i++ {
		v += "$" + string(i+1)
		k += "$" + string(i+1+len(keys))
		if i+1 != len(keys) {
			v += ","
			v += "'"
		}
	}
	return "insert into " + user + "(" + k + ") values (" + v + ");"
}

func registerUser(db *sql.DB, insertQuery, createQuery, uname, get, insert, email string) error {
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
	_, err = stmt.Exec(uname, get, insert, email)
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
// in his custom database.
func GetUsersData(db *sql.DB, uname, email string) []byte {
	row, err := db.Query("select email, get from users where uname = $1;", uname)
	if err != nil {
		return nil
	}
	defer row.Close()
	var e string
	var get string

	if row.Next() {
		row.Scan(&e, &get)
		if email != e {
			// wrong user
			return nil
		}
	} else {
		// wrong username
		return nil
	}

	rows, err := db.Query(get)
	data, err := readRows(rows)
	if err != nil {
		return nil
	}
	jsonInfo, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return jsonInfo
}

func readRows(rows *sql.Rows) ([][]Pair, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawResult := make([][]byte, len(cols))
	ret := [][]Pair{}
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		result := []Pair{}
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
