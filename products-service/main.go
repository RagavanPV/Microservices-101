package main

import (
    "products-service/configs"
    "products-service/routes"
    "github.com/gofiber/fiber/v2" 
)

func main() {
    app := fiber.New()

    //run database
    configs.ConnectDB()

    //routes
    routes.ProductRouter(app)

    app.Listen(":6000")
}