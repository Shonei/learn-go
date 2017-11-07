package main

import (
	"database/sql"
	"log"

	"github.com/Shonei/learn-go/distributed-systems/helpers"

	_ "github.com/lib/pq"
)

func main() {
	// http.HandleFunc("/", serve)
	// http.ListenAndServe(":8080", nil)

	connStr := "user=postgres dbname=ds sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// res := hlp.CreateTable([]byte(`{"User": "Shonei", "Keys": [{"Key":"one", "Value":""}, {"Key":"shdgd", "Value":""}, {"Key":"two", "Value":""}]}`), db, "fjghfg")
	// fmt.Println(res)
	hlp.GetUsersData(db, "Shonei", "fjghfg")
}
