package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/contact", handler2)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>welcome home</h1>\n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>contact us page</h1>\n")
}