package models

import (
    "time"
    "github.com/google/uuid"
)

type BudgetItem struct {
    ID                 uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    ProjectID          uuid.UUID  `gorm:"type:uuid" json:"project_id"`
    NoSP2K             *string    `json:"no_sp2k"`
    CostTypeID         uuid.UUID  `gorm:"type:uuid" json:"cost_type_id"`
    JenisBiayaName     *string    `json:"jenis_biaya_name"`
    KategoriAnggaran   *string    `json:"kategori_anggaran"`
    TotalAnggaran      *float64   `gorm:"type:numeric(20,2)" json:"total_anggaran"`
    DeskripsiAnggaran  *string    `json:"deskripsi_anggaran"`
    PeriodeBulan       *string    `json:"periode_bulan"`
    JumlahAnggaran     float64    `gorm:"type:numeric(20,2)" json:"jumlah_anggaran"`
    BulanKe            *int       `json:"bulan_ke"`
    ParentBudgetID     *uuid.UUID `gorm:"type:uuid" json:"parent_budget_id"`
    IsParent           bool       `gorm:"default:false" json:"is_parent"`
    CreatedAt          time.Time  `json:"created_date"`
    UpdatedAt          time.Time  `json:"updated_date"`
}

func (BudgetItem) TableName() string {
    return "budget_items"
}
