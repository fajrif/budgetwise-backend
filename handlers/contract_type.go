package handlers

import (
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
)

func GetContractTypes(c *fiber.Ctx) error {
    var contractTypes []models.ContractType
    if err := config.DB.Order("name ASC").Find(&contractTypes).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch contract types",
        })
    }

    return c.JSON(fiber.Map{
        "contract_types": contractTypes,
    })
}

func GetContractType(c *fiber.Ctx) error {
    id := c.Params("id")

    var contractType models.ContractType
    if err := config.DB.Where("id = ?", id).First(&contractType).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Contract type not found",
        })
    }

    return c.JSON(contractType)
}

func CreateContractType(c *fiber.Ctx) error {
    var contractType models.ContractType
    if err := c.BodyParser(&contractType); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Create(&contractType).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create contract type",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(contractType)
}

func UpdateContractType(c *fiber.Ctx) error {
    id := c.Params("id")

    var contractType models.ContractType
    if err := config.DB.Where("id = ?", id).First(&contractType).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Contract type not found",
        })
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&contractType).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update contract type",
        })
    }

    return c.JSON(contractType)
}

func DeleteContractType(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.ContractType{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete contract type",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Contract type not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Contract type deleted successfully",
    })
}
