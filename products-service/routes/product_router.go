package routes

import (
    "github.com/gofiber/fiber/v2"
    "products-service/services"
)


func ProductRouter(app *fiber.App) {
    app.Get("/product/:productId", services.GetAProduct)
    app.Get("/products", services.GetAllProducts)
}