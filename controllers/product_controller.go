package controllers

import (
    "ecommerce-backend/services"
    "ecommerce-backend/models"
    "net/http"
    "encoding/json"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := services.GetProducts()
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    json.NewDecoder(r.Body).Decode(&product)
    err := services.AddProduct(product)
    if err != nil {
        http.Error(w, "Error creating product", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}
