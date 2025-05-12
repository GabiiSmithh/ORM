package tests

import (
    "fmt"
    "go-mongo-orm/orm"
)

func QueryTest() {
    email := "anellyko@gmail.com"
    user, err := orm.FindUserByEmail(email)
    if err != nil {
        fmt.Println("Usuário não encontrado:", err)
        return
    }

    fmt.Printf("Usuário encontrado: %+v\n", user)
}
