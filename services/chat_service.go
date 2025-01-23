package services

import (
	"ecommerce-backend/models"
	"ecommerce-backend/repositories"
)

func SendMessage(message models.Message) error {
	return repositories.SaveMessage(message)
}

func FetchMessages() ([]models.Message, error) {
	return repositories.GetAllMessages()
}
