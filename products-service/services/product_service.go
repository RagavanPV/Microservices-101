package services

import (
    "context"
    "products-service/configs"
    "products-service/models"
    "products-service/responses"
    "net/http"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validate = validator.New()


func GetAllProducts(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

    return c.Status(http.StatusOK).JSON(
        responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": products}},
    )
}

func GetAProduct(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    productId := c.Params("productId")
    var product models.Product
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(productId)

    err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&product)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    return c.Status(http.StatusOK).JSON(responses.ProductResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": product}})
}