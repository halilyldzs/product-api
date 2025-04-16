package models

import (
	"time"
	
	"gorm.io/gorm"
)

// Product represents the structure for a product
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`           // Unique identifier
	Name        string         `json:"name" gorm:"not null"`           // Name of the product
	Description string         `json:"description"`                     // Detailed description
	Price       float64        `json:"price" gorm:"not null"`          // Price of the product
	Quantity    int            `json:"quantity" gorm:"default:0"`      // Quantity of the product
	CreatedAt   *time.Time     `json:"created_at,omitempty"`           // Timestamp of creation
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`           // Timestamp of last update
	SKU         string         `json:"sku"`                            // Stock Keeping Unit
	Barcode     string         `json:"barcode"`                        // Barcode of the product
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`                 // For soft delete support
	

	// Add other relevant fields like SKU, Quantity, CreatedAt, UpdatedAt etc. if needed
} 