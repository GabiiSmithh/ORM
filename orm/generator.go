package orm

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go-mongo-orm/config"
)

// Criação de coleções e índices únicos
func EnsureIndexes() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    db := config.Client.Database("orm_example")

    // Índice único em CPF (pessoas)
    pessoas := db.Collection("pessoas")
    pessoaIndex := mongo.IndexModel{ // Cria o modelo do índice
        Keys: bson.D{{Key: "cpf", Value: 1}},
        Options: options.Index().SetUnique(true), // Define o índice como único
    }

    // Verificação de erro
    _, err := pessoas.Indexes().CreateOne(ctx, pessoaIndex)
    if err != nil {
        panic("Erro ao criar índice em pessoas: " + err.Error())
    }

    // Índice único em ISBN (livros)
    livros := db.Collection("livros")
    livroIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "isbn", Value: 1}},
        Options: options.Index().SetUnique(true),
    }
    _, err = livros.Indexes().CreateOne(ctx, livroIndex)
    if err != nil {
        panic("Erro ao criar índice em livros: " + err.Error())
    }

    // Índice único em código (produtos)
    produtos := db.Collection("produtos")
    produtoIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "codigo", Value: 1}},
        Options: options.Index().SetUnique(true),
    }
    _, err = produtos.Indexes().CreateOne(ctx, produtoIndex)
    if err != nil {
        panic("Erro ao criar índice em produtos: " + err.Error())
    }
}
