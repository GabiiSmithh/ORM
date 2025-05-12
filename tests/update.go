package tests

import (
    "fmt"
    "go-mongo-orm/orm"
    "go.mongodb.org/mongo-driver/bson"
)

func UpdateTest() {
    email := "anellyko@gmail.com"
    update := bson.M{
        "age": 21,
    }

    result, err := orm.UpdateUserByEmail(email, update)
    if err != nil {
        fmt.Println("Erro ao atualizar:", err)
        return
    }

    fmt.Printf("Documentos modificados: %d\n", result.ModifiedCount)
}
