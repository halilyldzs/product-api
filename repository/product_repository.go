package repository

import (
	"errors"
	"time"
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
	// Set timestamps
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	
	query := `INSERT INTO products (name, description, price, quantity, created_at, updated_at, sku) 
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	// Execute the query
	result, err := database.DB.Exec(query, 
		product.Name, 
		product.Description, 
		product.Price,
		product.Quantity,
		product.CreatedAt,
		product.UpdatedAt,
		product.SKU)
	if err != nil {
		return err
	}
	
	// Get the auto-generated ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	// Set the ID in the product
	product.ID = uint(id)
	return nil
}

// GetAll retrieves all products from the database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `SELECT id, name, description, price, quantity, created_at, updated_at, sku FROM products`
	
	// Execute the query
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	// Parse the results
	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID, 
			&product.Name, 
			&product.Description, 
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.SKU); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return products, nil
}

// GetByID retrieves a product by its ID
func (r *ProductRepository) GetByID(id uint) (models.Product, error) {
	query := `SELECT id, name, description, price, quantity, created_at, updated_at, sku 
	          FROM products WHERE id = ?`
	
	var product models.Product
	err := database.DB.QueryRow(query, id).Scan(
		&product.ID, 
		&product.Name, 
		&product.Description, 
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.SKU)
	if err != nil {
		return models.Product{}, err
	}
	
	return product, nil
}

// Update updates a product in the database
func (r *ProductRepository) Update(product models.Product) error {
	// Update timestamp
	product.UpdatedAt = time.Now()
	
	query := `UPDATE products 
	          SET name = ?, description = ?, price = ?, quantity = ?, updated_at = ?, sku = ? 
	          WHERE id = ?`
	
	result, err := database.DB.Exec(query, 
		product.Name, 
		product.Description, 
		product.Price,
		product.Quantity,
		product.UpdatedAt,
		product.SKU,
		product.ID)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	
	return nil
}

// Delete removes a product from the database
func (r *ProductRepository) Delete(id uint) error {
	query := `DELETE FROM products WHERE id = ?`
	
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	
	return nil
} 