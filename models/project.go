package models

import (
    "time"
    "github.com/google/uuid"
)

type Project struct {
    ID                        uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    NoSP2K                    string       `gorm:"column:no_sp2k;not null" json:"no_sp2k"`
    NoPerjanjian              *string      `json:"no_perjanjian"`
    NoAmandemen               *string      `json:"no_amandemen"`
    TanggalPerjanjian         *time.Time   `json:"tanggal_perjanjian"`
    JudulPekerjaan            string       `gorm:"not null" json:"judul_pekerjaan"`
    JangkaWaktu               *int         `json:"jangka_waktu"`
    TanggalMulai              time.Time    `gorm:"not null" json:"tanggal_mulai"`
    TanggalSelesai            *time.Time   `json:"tanggal_selesai"`
    NilaiPekerjaan            float64      `gorm:"type:numeric(20,2);not null" json:"nilai_pekerjaan"`
    ManagementFee             *float64     `gorm:"type:numeric(20,2)" json:"management_fee"`
    TarifManagementFeePersen  *float64     `gorm:"type:numeric(5,2)" json:"tarif_management_fee_persen"`
    ClientID                  uuid.UUID    `gorm:"type:uuid" json:"client_id"`
		Client                    Client       `gorm:"foreignKey:ClientID" json:"client_details"`
    ContractTypeID            uuid.UUID    `gorm:"type:uuid" json:"contract_type_id"`
		ContractType              ContractType `gorm:"foreignKey:ContractTypeID" json:"contract_type_details"`
    StatusKontrak             string       `gorm:"default:'Active'" json:"status_kontrak"`
    CreatedBy                 string       `json:"created_by"`
    CreatedAt                 time.Time    `json:"created_date"`
    UpdatedAt                 time.Time    `json:"updated_date"`
}

func (Project) TableName() string {
    return "projects"
}
