package routes

import (
    "budgetwise-backend/handlers"
    "budgetwise-backend/middleware"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")

    // Public routes
    auth := api.Group("/auth")
    auth.Post("/register", handlers.Register)
    auth.Post("/login", handlers.Login)

    // Protected routes
    api.Use(middleware.AuthRequired())

    // User routes
    api.Get("/me", handlers.GetMe)
    api.Put("/me", handlers.UpdateMe)

    // Project routes
    projects := api.Group("/projects")
    projects.Get("/", handlers.GetProjects)
    projects.Get("/:id", handlers.GetProject)
    projects.Post("/", handlers.CreateProject)
    projects.Put("/:id", handlers.UpdateProject)
    projects.Delete("/:id", handlers.DeleteProject)

    // Budget Item routes
    budgets := api.Group("/budget-items")
    budgets.Get("/", handlers.GetBudgetItems)
    budgets.Get("/:id", handlers.GetBudgetItem)
    budgets.Post("/", handlers.CreateBudgetItem)
    budgets.Put("/:id", handlers.UpdateBudgetItem)
    budgets.Delete("/:id", handlers.DeleteBudgetItem)

    // Transaction routes
    transactions := api.Group("/transactions")
    transactions.Get("/", handlers.GetTransactions)
    transactions.Get("/:id", handlers.GetTransaction)
    transactions.Post("/", handlers.CreateTransaction)
    transactions.Put("/:id", handlers.UpdateTransaction)
    transactions.Delete("/:id", handlers.DeleteTransaction)

    // Cost Type routes (Admin only)
    costTypes := api.Group("/cost-types")
    costTypes.Get("/", handlers.GetCostTypes)
    costTypes.Get("/:id", handlers.GetCostType)
    costTypes.Post("/", middleware.AdminOnly(), handlers.CreateCostType)
    costTypes.Put("/:id", middleware.AdminOnly(), handlers.UpdateCostType)
    costTypes.Delete("/:id", middleware.AdminOnly(), handlers.DeleteCostType)
}
