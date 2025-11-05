package models

import (
    "time"
    "github.com/google/uuid"
)

type Transaction struct {
    ID                      uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    ProjectID               uuid.UUID  `gorm:"type:uuid" json:"project_id"`
    NoSP2K                  *string    `json:"no_sp2k"`
    TanggalTransaksi        time.Time  `gorm:"not null" json:"tanggal_transaksi"`
    TanggalPOTagihan        *time.Time `json:"tanggal_po_tagihan"`
    BulanRealisasi          *string    `json:"bulan_realisasi"`
    CostTypeID              uuid.UUID  `gorm:"type:uuid" json:"cost_type_id"`
    JenisBiayaName          *string    `json:"jenis_biaya_name"`
    DeskripsiRealisasi      *string    `json:"deskripsi_realisasi"`
    JumlahRealisasi         float64    `gorm:"type:numeric(20,2);not null" json:"jumlah_realisasi"`
    PersentaseManagementFee *float64   `gorm:"type:numeric(5,2)" json:"persentase_management_fee"`
    NilaiManagementFee      *float64   `gorm:"type:numeric(20,2)" json:"nilai_management_fee"`
    JumlahTenagaKerja       *int       `json:"jumlah_tenaga_kerja"`
    BuktiTransaksiURL       *string    `json:"bukti_transaksi_url"`
    CreatedBy               string     `json:"created_by"`
    CreatedAt               time.Time  `json:"created_date"`
    UpdatedAt               time.Time  `json:"updated_date"`
}

func (Transaction) TableName() string {
    return "transactions"
}
