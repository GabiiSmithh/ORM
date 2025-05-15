package orm

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go-mongo-orm/config"
)

// Criação automática de índices no MongoDB
func EnsureIndexes() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    users := config.Client.Database("orm_example").Collection("users")

    // Cria índice único no campo "email"
    indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "email", Value: 1}}, // Índice sobre o campo email
        Options: options.Index().SetUnique(true),  // Define como único
    }

    // Aplica no Mongo
    _, err := users.Indexes().CreateOne(ctx, indexModel)
    if err != nil {
        panic("Erro ao criar índice: " + err.Error())
    }
}
