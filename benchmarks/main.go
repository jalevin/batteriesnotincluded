package main

import (
	"fmt"
	"log"
	"net/http"
	//"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()                 // parse arguments, you have to call this by yourself
	//fmt.Println("form: ", r.Form) // print form information in server side
	//fmt.Println("path:", r.URL.Path)
	//fmt.Println("scheme:", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//fmt.Println("key:", k, "/ val:", strings.Join(v, "  "))
	//}
	fmt.Fprintf(w, "Hello world!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":8082", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
