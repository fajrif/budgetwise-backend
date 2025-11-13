package handlers

import (
		// "fmt"
    "budgetwise-backend/config"
    "budgetwise-backend/models"
    "github.com/gofiber/fiber/v2"
    // "github.com/google/uuid"
)

func GetProjects(c *fiber.Ctx) error {
    var projects []models.Project
    query := config.DB.Order("created_at DESC")

    // Filter by status if provided
    if status := c.Query("status_kontrak"); status != "" {
        query = query.Where("status_kontrak = ?", status)
    }

    if err := query.Find(&projects).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch projects",
        })
    }

		// Always return consistent JSON shape
    return c.JSON(fiber.Map{
        "projects": projects,
    })
}

func GetProject(c *fiber.Ctx) error {
    id := c.Params("id")

    var project models.Project
    if err := config.DB.Where("id = ?", id).First(&project).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Project not found",
        })
    }

    return c.JSON(fiber.Map{
        "project": project,
    })
}

func CreateProject(c *fiber.Ctx) error {
    var project models.Project
    if err := c.BodyParser(&project); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    userEmail := c.Locals("userEmail").(string)
    project.CreatedBy = userEmail

		// fmt.Printf("Project object (%%+v): %+v\n", project)

    if err := config.DB.Create(&project).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create project",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {
    id := c.Params("id")

    var project models.Project
    if err := config.DB.Where("id = ?", id).First(&project).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Project not found",
        })
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := config.DB.Model(&project).Updates(updateData).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update project",
        })
    }

    return c.JSON(project)
}

func DeleteProject(c *fiber.Ctx) error {
    id := c.Params("id")

    result := config.DB.Delete(&models.Project{}, "id = ?", id)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete project",
        })
    }

    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Project not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Project deleted successfully",
    })
}
