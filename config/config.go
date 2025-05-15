package config

import (
	"context" // controle de tempo de execução
	"log"     // registrar erros
	"time"    // duração da execução

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client //Cliente de conexão com o MongoDB

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // criação do contexto para 10s
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // conexão com o MongoDB na porta padrão 27017

	client, err := mongo.Connect(ctx, clientOptions) 
	if err != nil {                                  // verificação de erro
		log.Fatal(err)
	}

	Client = client // atribuição do cliente
}
