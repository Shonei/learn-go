package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/Shonei/learn-go/distributed-systems/helpers"

	_ "github.com/lib/pq"
)

func main() {
	// dsn := "host=shonei-portfolio:europe-west1:distributed-systems dbname=postgres user=postgres password=rz9h19cg sslmode=disable"
	// dsn := os.Getenv("POSTGRES_CONNECTION")

	db, err := sql.Open("cloudsqlpostgres", "host=shonei-portfolio:europe-west1:distributed-systems dbname=postgres user=postgres password=rz9h19cg sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", hlp.Root(db))
	http.HandleFunc("/get/", hlp.Get(db))
	http.HandleFunc("/post/", hlp.Post(db))
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.ListenAndServe(":8081", nil)

	// appengine.Main()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// registering a user
// curl --header "Email: teodor.shyl@gmail.com" --data '{"User": "Shyl", "Keys": [{"Key":"one"}, {"Key":"two"}]}' https://shonei-portfolio.appspot.com

// put data innto table
// curl --data '[{"Key": "one", "Value": "sdadgg"}, {"Key": "two", "Value": "sdgzvasok"}]' https://shonei-portfolio.appspot.com/post/Shyl

// get data from table
// curl https://shonei-portfolio.appspot.com/get/Shyl

//  curl --header "Email: teodor.shyl@gmail.com" --data '{"User": "tweets", "Keys": [{"Key":"uname"}, {"Key":"original"}, {"Key":"translated"}]}' https://shonei-portfolio.appspot.com
