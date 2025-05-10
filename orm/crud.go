package orm

import (
    "context"
    "go-mongo-orm/config"
    "go-mongo-orm/schema"
    "go.mongodb.org/mongo-driver/bson" // filtros/queries
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

func getCollection(name string) *mongo.Collection { // função para obter a coleção (danco = orm_exemple)
    return config.Client.Database("orm_example").Collection(name) 
}

// CREATE
func InsertUser(u schema.UserSchema) (*mongo.InsertOneResult, error) { // função para inserir um documento UserSchema na coleção users
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return getCollection("users").InsertOne(ctx, u) // retorna ID do novo documento ou erro.
}

// READ
func FindUserByEmail(email string) (schema.UserSchema, error) { // função para encontrar um documento UserSchema na coleção users
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user schema.UserSchema
    err := getCollection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user) // mapeia a struct com as chaves do documento e converte
    return user, err // retorna o usuario ou erro
}

// UPDATE
func UpdateUserByEmail(email string, update bson.M) (*mongo.UpdateResult, error) { // função que atualiza os campos do documento
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"email": email} // busca pelo campo email
    updateData := bson.M{"$set": update} // mapeamento dos campos a serem modificados (sem sobrescrever)

    return getCollection("users").UpdateOne(ctx, filter, updateData) // retorna o numero de documentos modificados ou erro
}

// DELETE
func DeleteUserByEmail(email string) (*mongo.DeleteResult, error) { // função que deleta o documento a partir de um email
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    return getCollection("users").DeleteOne(ctx, bson.M{"email": email}) // retorna o numero de documentos deletados ou erro
}

