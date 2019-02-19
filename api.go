package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	time2 "time"
)

type Time struct {
	Now string `json:"now,omitempty"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	time := Time{
		Now: time2.Now().String(),
	}

	fmt.Printf("Get request from %v", r.RemoteAddr)

	//w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(time)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}