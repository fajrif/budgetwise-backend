package handlers

import (
	"errors"
  "budgetwise-backend/config"
  "budgetwise-backend/models"
  "budgetwise-backend/helpers"
  "github.com/gofiber/fiber/v2"
)

// calculateDataTransaction adalah fungsi internal untuk mengisi field yang akan dikalkulasi
func calculateDataTransaction(tx *models.Transaction) error {
    // 1. Ambil data Project untuk mendapatkan persentase
    var project models.Project
    // GORM First(&project) akan mencari project berdasarkan ProjectID di tx
    if err := config.DB.Select("tarif_management_fee_persen").First(&project, "id = ?", tx.ProjectID).Error; err != nil {
        return errors.New("Project not found when calculating transaction data")
    }

    // 2. Hitung Bulan Realisasi (format "MMYYYY")
    monthYearStr := tx.TanggalTransaksi.Format("012006")
    tx.BulanRealisasi = &monthYearStr

    // 3. Hitung Nilai Management Fee
    if project.TarifManagementFeePersen != nil {
        percentage := *project.TarifManagementFeePersen / 100
        calculatedFee := tx.JumlahRealisasi * percentage
        tx.NilaiManagementFee = &calculatedFee
    } else {
        // Jika persentase di project null/kosong
        tx.NilaiManagementFee = nil
    }

    return nil
}

func GetTransactions(c *fiber.Ctx) error {
    var transactions []models.Transaction
    query := config.DB.Order("tanggal_transaksi DESC")

		// Filter Pencarian search
    if searchQuery := c.Query("search"); searchQuery != "" {
        // Wrap string pencarian dengan wildcard SQL LIKE
        searchTerm := "%" + searchQuery + "%"

        query = query.
            Joins("JOIN projects ON transactions.project_id = projects.id").
            Where("transactions.deskripsi_realisasi ILIKE ? OR projects.judul_pekerjaan ILIKE ? OR projects.no_sp2k ILIKE ?", searchTerm, searchTerm, searchTerm)
    }

    // Filter by project
    if projectID := c.Query("project_id"); projectID != "" {
        query = query.Where("project_id = ?", projectID)
    }

    // Filter by cost_type
    if costTypeID := c.Query("cost_type_id"); costTypeID != "" {
        query = query.Where("cost_type_id = ?", costTypeID)
    }

    // Filter by month
    if month := c.Query("bulan_realisasi"); month != "" {
        query = query.Where("bulan_realisasi = ?", month)
    }

    if err := query.
							Preload("Project").
							Preload("CostType").
							Find(&transactions).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch transactions",
        })
    }

    return c.JSON(fiber.Map{
        "transactions": transactions,
    })
}

func GetTransaction(c *fiber.Ctx) error {
    id := c.Params("id")
    query := config.DB

    var transaction models.Transaction
    if err := query.
							Preload("Project").
							Preload("CostType").
							Where("id = ?", id).First(&transaction).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Transaction not found",
        })
    }

    return c.JSON(transaction)
}

func CreateTransaction(c *fiber.Ctx) error {
    var transaction models.Transaction
    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := helpers.ValidateTransactionData(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

		// kalkulasi management fee
    if err := calculateDataTransaction(&transaction); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    userEmail := c.Locals("userEmail").(string)
    transaction.CreatedBy = userEmail

    if err := config.DB.Create(&transaction).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create transaction",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(transaction)
}

func UpdateTransaction(c *fiber.Ctx) error {
    id := c.Params("id")

    var transaction models.Transaction
    if err := config.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Transaction not found",
        })
    }

    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := helpers.ValidateTransactionData(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

		// kalkulasi management fee
    if err := calculateDataTransaction(&transaction); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    if err := config.DB.Model(&transaction).Updates(transaction).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update transaction",
        })
    }

    return c.JSON(transaction)
}

func DeleteTransaction(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.Transaction{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete transaction",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Transaction not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Transaction deleted successfully",
    })
}
