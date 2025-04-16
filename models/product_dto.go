package models

// ProductCreateDTO is used for creating a new product
// It includes only the fields that should be provided by the client
type ProductCreateDTO struct {
	Name        string  `json:"name"`         // Name of the product
	Description string  `json:"description"`  // Detailed description
	Price       float64 `json:"price"`        // Price of the product
	Quantity    int     `json:"quantity"`     // Quantity of the product
	SKU         string  `json:"sku"`          // Stock Keeping Unit
	Barcode     string  `json:"barcode"`      // Barcode of the product
}

// ToProduct converts a ProductCreateDTO to a Product entity
func (dto *ProductCreateDTO) ToProduct() Product {
	return Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Quantity:    dto.Quantity,
		SKU:         dto.SKU,
		Barcode:     dto.Barcode,
		// ID, CreatedAt and UpdatedAt will be set by the repository
	}
}

// ProductUpdateDTO is used for updating an existing product
// It's similar to ProductCreateDTO but will be applied to an existing Product
type ProductUpdateDTO struct {
	Name        string  `json:"name"`         // Name of the product
	Description string  `json:"description"`  // Detailed description
	Price       float64 `json:"price"`        // Price of the product
	Quantity    int     `json:"quantity"`     // Quantity of the product
	SKU         string  `json:"sku"`          // Stock Keeping Unit
	Barcode     string  `json:"barcode"`      // Barcode of the product
}

// ApplyToProduct updates an existing Product with the DTO values
func (dto *ProductUpdateDTO) ApplyToProduct(product *Product) {
	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Quantity = dto.Quantity
	product.SKU = dto.SKU
	product.Barcode = dto.Barcode
	// ID remains unchanged, UpdatedAt will be set by the repository
} 