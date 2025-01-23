package models

import "time"

type Message struct {
    ID        uint      `gorm:"primaryKey"`
    Content   string    `json:"content"`
    Sender    string    `json:"sender"`
    Timestamp time.Time `json:"timestamp"`
}
