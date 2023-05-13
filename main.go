package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestIn struct {
	Key   string `json:'key'`
	Value int    `json:'value'`
}

type RequestOut struct {
	Value int `json:'value'`
}

var RequestMap = make(map[string]int)

func WorkWithRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req RequestIn
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value := RequestMap[req.Key] + req.Value
	RequestMap[req.Key] = value
	resp := RequestOut{Value: value}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	port := "8080"
	http.HandleFunc("/", WorkWithRequest)
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
