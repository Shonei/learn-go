package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type Auth struct {
	Type  string `json:"token_type"`
	Token string `json:"access_token"`
}

func main() {
	http.HandleFunc("/search/", search)
	http.HandleFunc("/user/", user)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.ListenAndServe(":8080", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// make sure this is the expected route
	if !strings.Contains(r.URL.Path, "/user/") {
		http.Error(w, "Path doesnt' match", http.StatusBadRequest)
		return
	}

	query := strings.Trim(r.URL.Path, "/user/")

	url := "https://api.twitter.com/1.1/statuses/user_timeline.json?count=10&screen_name="
	res, err := getTweets(url, query)
	if err != nil {
		http.Error(w, "We were unable to retrieve tweets.", http.StatusInternalServerError)
	}

	w.Write(res)
}

func search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// make sure this is the expected route
	if !strings.Contains(r.URL.Path, "/search/") {
		http.Error(w, "Path doesnt' match", http.StatusBadRequest)
		return
	}

	query := strings.Trim(r.URL.Path, "/search/")

	url := "https://api.twitter.com/1.1/search/tweets.json?count=10&result_type=mixed&lang=en&q="
	res, err := getTweets(url, query)
	if err != nil {
		http.Error(w, "We were unable to retrieve tweets.", http.StatusInternalServerError)
	}

	w.Write(res)
}

func getTweets(url, search string) ([]byte, error) {
	authURL := "https://api.twitter.com/oauth2/token"
	body := []byte("grant_type=client_credentials")

	// prepare the request for auth token
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	appAuths := "1JbhCDuLRaYMg42SNfHOLDfyA:9FBlbBguoohxw4zmJwmF0DvXJKVTT3npdqPKt8d8I3QYkZqqzc"

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(appAuths)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8.")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read the app auth token
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	auth := Auth{}
	json.Unmarshal(resBody, &auth)

	// prepare the search request
	get, err := http.NewRequest("GET", url+search, nil)
	if err != nil {
		return nil, err
	}
	get.Header.Set("Authorization", "Bearer "+auth.Token)

	// get the response from the search
	tweets, err := http.DefaultClient.Do(get)
	if err != nil {
		return nil, err
	}

	// read the tweets as a []byte
	twits, err := ioutil.ReadAll(tweets.Body)
	if err != nil {
		return nil, err
	}

	return twits, nil
}
