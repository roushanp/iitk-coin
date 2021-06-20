package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roushanp/iitk-coin/database"
	"github.com/roushanp/iitk-coin/handlers"
)

func main() {

	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/secretpage", handlers.SecretPage)
	database.ConnectDB()
	
    fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
	
}