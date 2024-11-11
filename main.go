package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var processData = map[string]int{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handleReceiptPoints)

	addr := ":8080"

	split_address := strings.Split(addr, ":")

	var host = "0.0.0.0"

	if len(split_address[0]) > 0 {
		host = split_address[0]
	}

	port := split_address[1]

	fmt.Println("Starting server")
	fmt.Println(fmt.Sprintf("Listening on %s:%s", host, port))
	err := http.ListenAndServe(addr, r)

	if err != nil {
		fmt.Print("An error occurred starting the server")
		fmt.Println(err)
	}
}
