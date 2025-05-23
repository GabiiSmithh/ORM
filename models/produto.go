package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Produto struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nome        string             `bson:"nome"`
	Categoria   string             `bson:"categoria"`
	Preco       float64            `bson:"preco"`
	Disponivel  bool               `bson:"disponivel"`
	DataEntrada primitive.DateTime `bson:"data_entrada"`
}
