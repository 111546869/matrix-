package models

import "time"

type Todo struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"not null"`
    Title       string    `gorm:"not null"`
    Description string
    Status      string    `gorm:"default:'pending'"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}
