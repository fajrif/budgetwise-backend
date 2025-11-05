package handlers

import (
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
)

func GetBudgetItems(c *fiber.Ctx) error {
    var items []models.BudgetItem
    query := config.DB.Order("created_at DESC")

    // Filter by project
    if projectID := c.Query("project_id"); projectID != "" {
        query = query.Where("project_id = ?", projectID)
    }

    if err := query.Find(&items).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch budget items",
        })
    }

    return c.JSON(items)
}

func GetBudgetItem(c *fiber.Ctx) error {
    id := c.Params("id")

    var item models.BudgetItem
    if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Budget item not found",
        })
    }

    return c.JSON(item)
}

func CreateBudgetItem(c *fiber.Ctx) error {
    var item models.BudgetItem
    if err := c.BodyParser(&item); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Create(&item).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create budget item",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(item)
}

func UpdateBudgetItem(c *fiber.Ctx) error {
    id := c.Params("id")

    var item models.BudgetItem
    if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Budget item not found",
        })
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&item).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update budget item",
        })
    }

    return c.JSON(item)
}

func DeleteBudgetItem(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.BudgetItem{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete budget item",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Budget item not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Budget item deleted successfully",
    })
}
