package repository

import (
	"errors"
	"product-api/database"
	"product-api/models"
)

// ProductRepository handles database operations for products
type ProductRepository struct{}

// NewProductRepository creates a new product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// Create adds a new product to the database
func (r *ProductRepository) Create(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

// GetAll retrieves all products from the database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}

// GetByID retrieves a product by its ID
func (r *ProductRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}

// Update updates a product in the database
func (r *ProductRepository) Update(product models.Product) error {
	result := database.DB.Save(&product)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	
	return nil
}

// Delete removes a product from the database
func (r *ProductRepository) Delete(id uint) error {
	result := database.DB.Delete(&models.Product{}, id)
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	
	return nil
} 