package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
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