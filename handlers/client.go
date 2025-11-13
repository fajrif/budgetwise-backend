package handlers

import (
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
)

func GetClients(c *fiber.Ctx) error {
    var clients []models.Client
    query := config.DB.Order("created_at DESC")

    if err := query.Find(&clients).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch clients",
        })
    }

		// Always return consistent JSON shape
    return c.JSON(fiber.Map{
        "clients": clients,
    })
}

func GetClient(c *fiber.Ctx) error {
    id := c.Params("id")

    var client models.Client
    if err := config.DB.Where("id = ?", id).First(&client).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Client not found",
        })
    }

    return c.JSON(fiber.Map{
        "client": client,
    })
}

func CreateClient(c *fiber.Ctx) error {
    var client models.Client
    if err := c.BodyParser(&client); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Create(&client).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create client",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(client)
}

func UpdateClient(c *fiber.Ctx) error {
    id := c.Params("id")

    var client models.Client
    if err := config.DB.Where("id = ?", id).First(&client).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Client not found",
        })
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&client).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update client",
        })
    }

    return c.JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.Client{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete client",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Client not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Client deleted successfully",
    })
}
