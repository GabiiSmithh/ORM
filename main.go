package main

import (
    "go-mongo-orm/config"
    "go-mongo-orm/orm"
    "go-mongo-orm/tests"
)

func main() {
    // Conecta ao MongoDB
    config.Connect()

    // Garante que o índice único no campo "email" seja criado
    orm.EnsureIndexes()

    // Executa os casos de uso
    tests.InsertTest()
    tests.QueryTest()
    tests.UpdateTest()
}
