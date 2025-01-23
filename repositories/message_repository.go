package repositories

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
)

func SaveMessage(message models.Message) error {
	result := config.DB.Create(&message)
	return result.Error
}

func GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	result := config.DB.Find(&messages)
	return messages, result.Error
}
