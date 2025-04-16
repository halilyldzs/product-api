package main

import (
	"log"
	// "net/http" // Removed standard library http
	// "fmt" // Removed standard library fmt

	"github.com/gofiber/fiber/v2"
	"product-api/routes" // Import the routes package
)

// func helloHandler(w http.ResponseWriter, r *http.Request) { // Removed net/http handler
// 	fmt.Fprintf(w, "hello world")
// }

// @title Product API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupProductRoutes(app)
	routes.SetupSwaggerRoutes(app) // Setup Swagger routes

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
} 