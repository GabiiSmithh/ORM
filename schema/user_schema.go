package schema

import (
    "go.mongodb.org/mongo-driver/bson/primitive" // pacote para manipulação de BSON
    "go-mongo-orm/models"
)

type UserSchema struct { // tags mapeiam os campos da struct e os documentos do banco
    ID    primitive.ObjectID `bson:"_id,omitempty"`
    Name  string             `bson:"name"`
    Email string             `bson:"email"`
    Age   int                `bson:"age"`
} // name, email e age são exatamente como no banco, mas ID é um ObjectID pq permite ser ignorado na inserção

// Conversão do model.User (aplicação) para UserSchema (banco)
func FromModel(u model.User) UserSchema {
    return UserSchema{ // não tem ID, pois é gerado automaticamente pelo banco após a inserção
        Name:  u.Name,
        Email: u.Email,
        Age:   u.Age,
    }
}

// Conversão de UserSchema para model.User
func ToModel(u UserSchema) model.User {
    return model.User{
        ID:    u.ID.Hex(), // conversão de ObjectID para string
        Name:  u.Name,
        Email: u.Email,
        Age:   u.Age,
    }
}
