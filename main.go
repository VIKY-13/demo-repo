package main

import (
	// "fmt"
	"log"
	// "net/http"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Get("/env",func(c *fiber.Ctx) error {
		return c.SendString("env")
	})
	port :="3000"
	log.Fatal(app.Listen("0.0.0.0:"+port))
	
}

// func Hello(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w,"Hello")
// }

// func welcome(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w,"welcome")
// }