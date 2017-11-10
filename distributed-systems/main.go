package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Shonei/learn-go/distributed-systems/helpers"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres dbname=ds sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// fmt.Println("аогяожгяаого")
	http.HandleFunc("/", hlp.Root(db))
	http.HandleFunc("/get/", hlp.Get(db))
	http.HandleFunc("/post/", hlp.Post(db))
	http.ListenAndServe(":8080", nil)
}

// registering a user
// curl --header "Email: teodor.shyl@gmail.com" --data '{"User": "Shyl", "Keys": [{"Key":"one"}, {"Key":"two"}]}' http://localhost:8080

// put data innto table
// curl --data '[{"Key": "one", "Value": "sdadgg"}, {"Key": "two", "Value": "sdgzvasok"}]' http://localhost:8080/post/Shyl

// get data from table
// curl http://localhost:8080/get/Shyl
