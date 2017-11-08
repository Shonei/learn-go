package hlp

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

		var table Table
		// converts body to swrings and makes a reader to it
		// passes the reader to the json decoder and decodes the data to table
		json.NewDecoder(strings.NewReader(string(body))).Decode(&table)

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
		fmt.Println(r.URL.Path)
		// make sure we have the right path
		if !strings.Contains(r.URL.Path, "/get/") {
			http.Error(w, "Path doesnt' match", http.StatusBadRequest)
			return
		}

		uname := strings.Trim(r.URL.Path, "/get/")

		email := r.Header.Get("email")
		if email == "" {
			w.Write([]byte("Please prove an email in the email http header"))
			return
		}
		fmt.Println(uname)
		fmt.Println(email)
		data := GetUsersData(db, uname, email)
		if data == nil {
			w.Write([]byte("You dont have any data stored"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
