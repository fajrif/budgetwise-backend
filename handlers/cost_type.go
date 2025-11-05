package handlers

import (
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
)

func GetCostTypes(c *fiber.Ctx) error {
    var costTypes []models.CostType
    if err := config.DB.Order("nama_biaya ASC").Find(&costTypes).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch cost types",
        })
    }

    return c.JSON(costTypes)
}

func GetCostType(c *fiber.Ctx) error {
    id := c.Params("id")

    var costType models.CostType
    if err := config.DB.Where("id = ?", id).First(&costType).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Cost type not found",
        })
    }

    return c.JSON(costType)
}

func CreateCostType(c *fiber.Ctx) error {
    var costType models.CostType
    if err := c.BodyParser(&costType); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Create(&costType).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create cost type",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(costType)
}

func UpdateCostType(c *fiber.Ctx) error {
    id := c.Params("id")

    var costType models.CostType
    if err := config.DB.Where("id = ?", id).First(&costType).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Cost type not found",
        })
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&costType).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update cost type",
        })
    }

    return c.JSON(costType)
}

func DeleteCostType(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.CostType{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete cost type",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Cost type not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Cost type deleted successfully",
    })
}
