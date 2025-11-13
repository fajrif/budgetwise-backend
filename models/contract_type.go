package models

import (
    "time"
    "github.com/google/uuid"
)

type ContractType struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    Name        string    `gorm:"not null" json:"name"`
    Code        *string   `json:"code"`
    Description *string   `json:"description"`
    CreatedAt   time.Time `json:"created_date"`
    UpdatedAt   time.Time `json:"updated_date"`
}

func (ContractType) TableName() string {
    return "contract_types"
}
