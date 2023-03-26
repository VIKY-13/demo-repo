package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", Hello)
	http.HandleFunc("/welcome",welcome)
	http.ListenAndServe(":8080",nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello")
}

func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"welcome")
}