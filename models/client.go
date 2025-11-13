package models

import (
    "time"
    "github.com/google/uuid"
)

type Client struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    Name        string    `gorm:"not null" json:"name"`
    ContactName *string   `json:"contact_name"`
    Phone       *string   `json:"phone"`
    Address     *string   `json:"address"`
    CreatedAt   time.Time `json:"created_date"`
    UpdatedAt   time.Time `json:"updated_date"`
}

func (Client) TableName() string {
    return "clients"
}
