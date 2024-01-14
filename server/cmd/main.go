package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type M map[string]interface{}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(M{"response": "Hello World"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	log.Println("Server Started")
	srv := http.ListenAndServe(":8000", r)
	if srv != nil {
		log.Fatalf(srv.Error())
	}
}
