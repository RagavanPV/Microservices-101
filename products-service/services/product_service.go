package services

import (
    "context"
    "products-service/configs"
    "products-service/models"
    "products-service/responses"
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

    Logger "github.com/RagavanPV/Microservices-101/libraries/go_logger"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")


func GetAllProducts(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    logger := Logger.LoggerFromContext(ctx)
    logger.Info().Msg("Retrieving All Products")
    var products []models.Product
    defer cancel()

    results, err := productCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleProduct models.Product
        if err = results.Decode(&singleProduct); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }

        products = append(products, singleProduct)
    }
    logger.Info().Msg("Retrieved All Products")

    return c.Status(http.StatusOK).JSON(
        responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": products}},
    )
}

func GetAProduct(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    logger := Logger.LoggerFromContext(ctx)
    logger.Info().Msg("Retrieving A Product")
    productId := c.Params("productId")
    var product models.Product
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(productId)

    err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&product)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
    logger.Info().Msg("Retrieved A Product")
    return c.Status(http.StatusOK).JSON(responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": product}})
}