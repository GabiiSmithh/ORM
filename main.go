package main

import (
    "fmt"
    "go-mongo-orm/config"
    "go-mongo-orm/models"
    "go-mongo-orm/orm"
    "go-mongo-orm/schema"
)

func main() {
    config.Connect()
    orm.EnsureIndexes()

    user := model.User{
        Name:  "João",
        Email: "joao@silva.com",
        Age:   30,
    }

    // Inserção
    inserted, err := orm.InsertUser(schema.FromModel(user))
    if err != nil {
        panic(err)
    }
    fmt.Println("Usuário inserido com ID:", inserted.InsertedID)

    // Leitura
    found, err := orm.FindUserByEmail("joao@silva.com")
    if err != nil {
        panic(err)
    }
    fmt.Println("Usuário encontrado:", schema.ToModel(found))
}
