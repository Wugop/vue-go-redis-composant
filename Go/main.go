package main

import (
	handlers "go-redis/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	port := ":8085"
	r := mux.NewRouter()
	r.HandleFunc("/sha256", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		handlers.Sha256Handler(w, r)
	})
	r.HandleFunc("/sha256/{key}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		handlers.Sha256HandlerKey(w, r)
	})
	log.Fatal(http.ListenAndServe(port, r))
}
