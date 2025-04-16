package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	// "net/http" // Removed standard library http
	// "fmt" // Removed standard library fmt

	"github.com/gofiber/fiber/v2"
	"product-api/database"
	"product-api/routes" // Import the routes package
)

// func helloHandler(w http.ResponseWriter, r *http.Request) { // Removed net/http handler
// 	fmt.Fprintf(w, "hello world")
// }

// @title Product API
// @version 1.0
// @description This is a sample server for a product API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Create a new Fiber app with custom error handling
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return a generic error if something goes wrong
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup routes
	routes.SetupProductRoutes(app)
	routes.SetupSwaggerRoutes(app) // Setup Swagger routes

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
} 