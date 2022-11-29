package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go-redis/redisconn"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type my_struct struct {
	Key string
}

func Sha256HandlerKey(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		red := redisconn.GetRedisConnection()
		vars := mux.Vars(r)
		key, ok := vars["key"]
		if ok {
			response, err := red.Get(key).Result()
			if err != nil {
				log.Print(err.Error())
				http.Error(w, "Internal server error - "+err.Error(), 500)
				return
			}
			w.Write([]byte(response))
		}
	case "DELETE":
		red := redisconn.GetRedisConnection()
		vars := mux.Vars(r)
		key, ok := vars["key"]
		if ok {
			red.Del(key)
		}
	}
}

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		red := redisconn.GetRedisConnection()
		var ArrayResponse []string
		iter := red.Scan(0, "", 0).Iterator()
		for iter.Next() {
			ArrayResponse = append(ArrayResponse, iter.Val())
		}
		if err := iter.Err(); err != nil {
			panic(err)
		}
		var response = strings.Join(ArrayResponse, "-")
		w.Write([]byte(response))
	case "POST":
		var s my_struct
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Print(err)
			return
		}
		defer r.Body.Close()

		red := redisconn.GetRedisConnection()
		sum := sha256.Sum256([]byte(s.Key))
		sumString := fmt.Sprintf("%x", sum)
		_, err = red.SetNX(s.Key, sumString, 0).Result()
		if err != nil {
			http.Error(w, "Internal server error", 500)
			log.Print(err)
			return
		}
		w.Write([]byte("Done"))
	}
}
