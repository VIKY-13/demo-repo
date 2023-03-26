package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

)

func main(){
	os.Setenv("PORT","4000")
	http.HandleFunc("/", Hello)
	http.HandleFunc("/w", welcome)
	port :=os.Getenv("PORT")
	http.ListenAndServe("0.0.0.0:"+port,nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello")
}

func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"welcome")
}