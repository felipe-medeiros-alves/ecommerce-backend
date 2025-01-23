package controllers

import (
	"ecommerce-backend/clients"
	"net/http"
)

func ChatEndpoint(w http.ResponseWriter, r *http.Request) {
	clients.HandleConnections(w, r)
}
