package tests

import (
    "fmt"
    "go-mongo-orm/orm"
    "go-mongo-orm/schema"
)

func InsertTest() {
    user := schema.UserSchema{
        Name:  "Anelly",
        Email: "anellykov@gmail.com",
        Age:   20,
    }

    result, err := orm.InsertUser(user)
    if err != nil {
        fmt.Println("Erro ao inserir:", err)
        return
    }

    fmt.Println("Usuário inserido com ID:", result.InsertedID)
}
