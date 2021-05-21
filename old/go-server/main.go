package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("hello")
	defer r.Body.Close() // Example関数が終了する時に実行されるdeferステートメント
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Parameter struct {
		Msg string `json:"msg"`
	}
	var param Parameter

	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf(param.Msg)
}

func handleOption(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	log.Printf("hello")
	router := httprouter.New()
	router.OPTIONS("/", handleOption)
	router.POST("/", hello)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
