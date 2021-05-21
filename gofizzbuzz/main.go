package main

import (
	"encoding/json"
	"net/http"

	"github.com/lastrust/recruitment/gofizzbuzz/fizzbuzz"
)

func main() {
	http.HandleFunc("/fizzbuzz", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := fizzbuzz.Run()
	res := struct{ Result int `json: "result"`}{ result }
	json, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}