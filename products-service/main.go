package main

import (
    "context"
    "products-service/configs"
    "products-service/routes"
    "github.com/gofiber/fiber/v2" 
    Logger "github.com/ragavan/go_logger"
    LoggerModel "github.com/ragavan/go_logger/model"
)

func main() {
    app := fiber.New()
    ctx := context.Background()

    logger, err := Logger.Init(LoggerModel.LogOptions{Level: "info"})
    Logger.ContextWithLogger(ctx, logger)
    if err != nil {
		return 
	}
    logger.Info().Msg("Started Server")

    //run database
    configs.ConnectDB()

    //routes
    routes.ProductRouter(app)

    app.Listen(":6000")
}