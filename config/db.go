package config

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var client *mongo.Client // Representação da conexão com o banco de dados
var database *mongo.Database // Referência ao banco de dados

func Connect(uri string, dbName string) { // Função para conectar ao banco de dados (endereço e nome do banco)
    var err error
    client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri)) // Cria a conexão sem timeout
    if err != nil {
        log.Fatal(err)
    }
    database = client.Database(dbName) // Atribuição do banco de dados á uma variável.
}

func GetCollection(name string) *mongo.Collection { // Função para retornar a coleção do banco de dados
    return database.Collection(name)
}
