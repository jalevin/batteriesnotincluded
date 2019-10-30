package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!") // send data to client side
}
func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":8082", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
