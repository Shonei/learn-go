package hlp

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	D string
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	d := Data{D: "Welcome to go"}
	json.NewEncoder(w).Encode(d)
}
