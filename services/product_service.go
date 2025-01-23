package services

import (
    "ecommerce-backend/models"
    "ecommerce-backend/repositories"
)

func GetProducts() ([]models.Product, error) {
    return repositories.GetAllProducts()
}

func AddProduct(product models.Product) error {
    return repositories.CreateProduct(product)
}
