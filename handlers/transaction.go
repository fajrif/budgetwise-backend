package handlers

import (
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
)

func GetTransactions(c *fiber.Ctx) error {
    var transactions []models.Transaction
    query := config.DB.Order("tanggal_transaksi DESC")

    // Filter by project
    if projectID := c.Query("project_id"); projectID != "" {
        query = query.Where("project_id = ?", projectID)
    }

    // Filter by month
    if month := c.Query("bulan_realisasi"); month != "" {
        query = query.Where("bulan_realisasi = ?", month)
    }

    if err := query.Find(&transactions).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch transactions",
        })
    }

    return c.JSON(transactions)
}

func GetTransaction(c *fiber.Ctx) error {
    id := c.Params("id")

    var transaction models.Transaction
    if err := config.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
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

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&transaction).Updates(updateData).Error; err != nil {
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
