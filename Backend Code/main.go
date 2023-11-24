// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"your-package/config"
	"your-package/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Set up Gin router
	router := routes.SetupRouter()

	// Run the server
	port := 8080
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on http://localhost%s", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
