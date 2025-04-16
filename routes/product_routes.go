package routes

import (
	"product-api/models"
	"product-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ProductRepository is used for database operations
var productRepo *repository.ProductRepository

// init initializes the product repository
func init() {
	productRepo = repository.NewProductRepository()
}

// addProductHandler handles POST requests to create a new product
// @Summary Create a new product
// @Description Add a new product to the system
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.ProductCreateDTO true "Product to add"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func addProductHandler(c *fiber.Ctx) error {
	// Parse request body into ProductCreateDTO struct
	var productDTO models.ProductCreateDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate input
	if productDTO.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if productDTO.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Price must be greater than 0",
		})
	}

	if productDTO.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Quantity cannot be negative",
		})
	}

	// Convert DTO to Product entity
	product := productDTO.ToProduct()

	// Add to database
	err := productRepo.Create(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	// Return success response with the created product
	return c.Status(fiber.StatusCreated).JSON(product)
}

// getProductsHandler handles GET requests to retrieve all products
// @Summary Get all products
// @Description Retrieve all products from the system
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func getProductsHandler(c *fiber.Ctx) error {
	products, err := productRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(products)
}

// getProductHandler handles GET requests to retrieve a specific product
// @Summary Get a product by ID
// @Description Retrieve a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func getProductHandler(c *fiber.Ctx) error {
	// Get product ID from URL
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Get product from database
	product, err := productRepo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(product)
}

// updateProductHandler handles PUT requests to update a product
// @Summary Update a product
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductUpdateDTO true "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
func updateProductHandler(c *fiber.Ctx) error {
	// Get product ID from URL
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Parse request body
	var productDTO models.ProductUpdateDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate input
	if productDTO.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if productDTO.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Price must be greater than 0",
		})
	}

	if productDTO.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Quantity cannot be negative",
		})
	}

	// Get existing product
	product, err := productRepo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	// Apply changes from DTO
	productDTO.ApplyToProduct(&product)

	// Update product in database
	err = productRepo.Update(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.JSON(product)
}

// deleteProductHandler handles DELETE requests to remove a product
// @Summary Delete a product
// @Description Remove a product from the system
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func deleteProductHandler(c *fiber.Ctx) error {
	// Get product ID from URL
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	// Delete product from database
	err = productRepo.Delete(uint(id))
	if err != nil {
		if err.Error() == "product not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

// SetupProductRoutes configures the product related routes
func SetupProductRoutes(app *fiber.App) {
	// Group routes under /api/v1
	api := app.Group("/api/v1")

	// Product routes
	products := api.Group("/products")
	products.Post("/", addProductHandler)         // POST /api/v1/products
	products.Get("/", getProductsHandler)         // GET /api/v1/products
	products.Get("/:id", getProductHandler)       // GET /api/v1/products/:id
	products.Put("/:id", updateProductHandler)    // PUT /api/v1/products/:id
	products.Delete("/:id", deleteProductHandler) // DELETE /api/v1/products/:id
} 