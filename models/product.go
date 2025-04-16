package models

import "time"

// Product represents the structure for a product
type Product struct {
	ID          uint    `json:"id"`           // Unique identifier
	Name        string  `json:"name"`         // Name of the product
	Description string  `json:"description"`  // Detailed description
	Price       float64 `json:"price"`        // Price of the product
	Quantity    int     `json:"quantity"`     // Quantity of the product
	CreatedAt   time.Time `json:"created_at"` // Timestamp of creation
	UpdatedAt   time.Time `json:"updated_at"` // Timestamp of last update
	SKU         string    `json:"sku"`        // Stock Keeping Unit

	// Add other relevant fields like SKU, Quantity, CreatedAt, UpdatedAt etc. if needed
} 