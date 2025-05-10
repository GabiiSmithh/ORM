package orm

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "time"
)

func EnsureIndexes() { // função que realiza a migração e criação de índices
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    users := getCollection("users") // obtém a coleção de usuários

    // Criar índice único no email
    indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "email", Value: 1}}, // onde será o índice
        Options: options.Index().SetUnique(true), // define que será único
    }

    _, _ = users.Indexes().CreateOne(ctx, indexModel) // cria o índice e retorna o resultado ou erro
}
