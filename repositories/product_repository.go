package repositories

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"errors"
)

func GetAllProducts() ([]models.Product, error) {
	if config.DB == nil {
		return nil, errors.New("database connection is not established")
	}

	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product models.Product) error {
	if config.DB == nil {
		return errors.New("database connection is not established")
	}

	result := config.DB.Create(&product)
	return result.Error
}
