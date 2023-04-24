package main

import (
    "context"
    "os"
    "products-service/configs"
    "products-service/routes"
    "github.com/gofiber/fiber/v2" 
    Logger "github.com/RagavanPV/Microservices-101/libraries/go_logger"
    LoggerModel "github.com/RagavanPV/Microservices-101/libraries/go_logger/model"
)

func main() {
    app := fiber.New()
    ctx := context.Background()
    runLogFile, _ := os.OpenFile(
        "logs/products-service.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
    multi := Logger.MultiLevelWriter(os.Stdout, runLogFile)
    logger, err := Logger.Init(LoggerModel.LogOptions{Level: "info", Writer: multi})
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