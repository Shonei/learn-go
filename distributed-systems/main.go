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

	http.HandleFunc("/", hlp.Root(db))
	http.HandleFunc("/get/", hlp.Get(db))
	http.ListenAndServe(":8080", nil)

	// res := hlp.CreateTable([]byte(`{"User": "Shonei", "Keys": [{"Key":"one"}, {"Key":"shdgd"}, {"Key":"two"}]}`), db, "fjghfg")
	// fmt.Println(res)
	// hlp.GetUsersData(db, "Shonei", "fjghfg")
	// d := []hlp.Pair{hlp.Pair{Key: "one", Value: "sdadgg"},
	// 	hlp.Pair{Key: "shdgd", Value: "sdgereyh"},
	// 	hlp.Pair{Key: "two", Value: "sdgzvasok"}}
	// foo, err := json.Marshal(d)
	// if err != nil {
	// 	return
	// }
	// hlp.InsertIntoUser(db, foo, "Shonei")
}

// curl --header "Email: teodor.shyl@gmail.com" --data '{"User": "dfhdf", "Keys": [{"Key":"one"}, {"Key":"shdgd"}, {"Key":"two"}]}'
