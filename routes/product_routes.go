package routes

import (
	"product-api/models"
	"github.com/gofiber/fiber/v2"
)

// In-memory products store (temporary solution until a database is implemented)
var products = make([]models.Product, 0)
var nextID uint = 1

// addProductHandler handles POST requests to create a new product
// @Summary Create a new product
// @Description Add a new product to the system
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product to add"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func addProductHandler(c *fiber.Ctx) error {
	// Parse request body into Product struct
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate input
	if product.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if product.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Price must be greater than 0",
		})
	}

	// Set a unique ID for the product
	product.ID = nextID
	nextID++

	// Add to in-memory storage
	products = append(products, product)

	// Return success response with the created product
	return c.Status(fiber.StatusCreated).JSON(product)
}

// SetupProductRoutes configures the product related routes
func SetupProductRoutes(app *fiber.App) {
	// Group routes under /api/v1
	api := app.Group("/api/v1")

	// Product routes
	products := api.Group("/products")
	products.Post("/", addProductHandler) // POST /api/v1/products
	// Add other product routes here (GET, PUT, DELETE) if needed
} 