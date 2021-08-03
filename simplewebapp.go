package main

import(
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Simple Web App using Golang")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "About page")
}

func main(){
	http.HandleFunc("/", index_handler)//Handles functiosn that will correspond to paths
	http.HandleFunc("/about", about_handler)//About page on the web application
	http.ListenAndServe(":8000", nil)//Listen on port 8000, server is nil for now
}

//To run this: go run simplewebapp.go, and then open a browser, heading to 127.0.0.1:8000.
