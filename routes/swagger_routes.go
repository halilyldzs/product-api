package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "product-api/docs" // Import generated docs
)

// SetupSwaggerRoutes configures the Swagger UI route
// @title           Product API
// @version         1.0
// @description     This is a sample server for a product API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func SetupSwaggerRoutes(app *fiber.App) {
	// Route for Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Default route for redirecting to swagger if needed
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
	})
} 