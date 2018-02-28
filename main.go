// created by me
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"GoCasst/Auth"
	"GoCasst/Cass"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	if Auth.AuthCheck(w, r) {
		json.NewEncoder(w).Encode(ResponseStatus{Status: "OK", Code: 200})
		fmt.Println(Cass.LogTime() + ", Heartbeat Received with Status 200")
		return
	}
	json.NewEncoder(w).Encode(ResponseStatus{Status: "Unauthorized", Code: 401})
	fmt.Println(Cass.LogTime() + ", Heartbeat request - Unauthorized")
}

func main() {
	//var handler http.HandlerFunc
	CassSession := Cass.Session
	defer CassSession.Close()
	// router endpoints and handlers (endpoints.go)
	router := NewRouter()

	serv := &http.Server{
		Addr:           "127.0.0.1:8088",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 50,
	}

	log.Fatal(serv.ListenAndServe())
}
