package hlp

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Root returns a handler function
// it takes the database as an arguments so I can use it
// while I handle the request
func Root(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	// func Root() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// make sure this is the expected route
		if r.URL.Path != "/" {
			http.Error(w, "Path doesnt' match", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// Read json data and users email making sure he is loged in
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "No JSON provided", http.StatusBadRequest)
			return
		}

		email := r.Header.Get("email")
		if email == "" {
			w.Write([]byte("Please prove an email in the email http header"))
			return
		}

		// register the users
		res := CreateTable(body, db, email)

		json.NewEncoder(w).Encode(res)
	}
}

// Get returns the handler function for get requests
// It again wraps the handler around to give it access to the
// sql database.
func Get(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	// func Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// make sure we have the right path
		if !strings.Contains(r.URL.Path, "/get/") {
			http.Error(w, "Path doesnt' match", http.StatusBadRequest)
			return
		}

		uname := strings.TrimPrefix(r.URL.Path, "/get/")

		data := GetUsersData(db, uname)
		if data == nil {
			w.Write([]byte("You dont have any data stored"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// Post will handle request for users to put data into their table.
// It is passed a database and returns a http handler function.
func Post(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// make sure we have the right path
		if !strings.Contains(r.URL.Path, "/post/") {
			// fmt.Println("err")
			http.Error(w, "Path doesnt' match", http.StatusBadRequest)
			return
		}

		uname := strings.TrimPrefix(r.URL.Path, "/post/")

		// get the data to put into the db

		// j := &[]Pair{}
		// json.NewDecoder(r.Body).Decode(&j)
		// fmt.Println(j)

		body, err := ioutil.ReadAll(r.Body)
		// fmt.Println("Body")
		// fmt.Println(string(body))
		if err != nil {
			// fmt.Println(err)
			http.Error(w, "No JSON provided", http.StatusBadRequest)
			return
		}

		err = InsertIntoUser(db, body, uname)
		if err != nil {
			// fmt.Println(err)
			http.Error(w, "We were unable to update your data", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Update successful"))
	}
}
