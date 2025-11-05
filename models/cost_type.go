package models

import (
    "time"
    "github.com/google/uuid"
)

type CostType struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    NamaBiaya   string    `gorm:"not null" json:"nama_biaya"`
    Kode        *string   `json:"kode"`
    Deskripsi   *string   `json:"deskripsi"`
    CreatedAt   time.Time `json:"created_date"`
    UpdatedAt   time.Time `json:"updated_date"`
}

func (CostType) TableName() string {
    return "cost_types"
}
