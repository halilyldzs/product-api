package models

// Product represents the structure for a product
type Product struct {
	ID          uint    `json:"id"`           // Unique identifier
	Name        string  `json:"name"`         // Name of the product
	Description string  `json:"description"`  // Detailed description
	Price       float64 `json:"price"`        // Price of the product
	// Add other relevant fields like SKU, Quantity, CreatedAt, UpdatedAt etc. if needed
} 