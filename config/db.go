package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Conectar(uri string) (*mongo.Client, error) {
	cliente, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri)) // Inicia o cliente MongoDB
	if err != nil {
		return nil, err
	}
	if err = cliente.Ping(context.Background(), nil); err != nil { // Ping para testar a conexão
		return nil, err
	}
	return cliente, nil
}
