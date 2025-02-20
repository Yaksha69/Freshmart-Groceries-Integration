// main.go
package main

import (
	"employee-service/routes"
	"log"
	"net/http"
)

func main() {
	// Register routes
	routes.RegisterRoutes()

	// Start the server
	log.Println("Server started on :3004")
	err := http.ListenAndServe(":3004", nil)
	if err != nil {
		log.Fatal(err)
	}
}
