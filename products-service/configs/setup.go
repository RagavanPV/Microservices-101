package configs

import (
    "context"
    Logger "github.com/RagavanPV/Microservices-101/libraries/go_logger"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client  {
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    logger := Logger.LoggerFromContext(ctx)
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        log.Fatal(err)
    }

    
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    logger.Info().Msg("Connected to MongoDB")
    return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database(EnvMongoDatabase()).Collection(collectionName)
    return collection
}