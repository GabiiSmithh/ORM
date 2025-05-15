package orm

import (
    "context"
    "go-mongo-orm/config"
    "reflect"
    "strings"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// obtém o nome da coleção a partir do nome da struct
func getCollectionName(doc interface{}) string {
    t := reflect.TypeOf(doc)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    return strings.ToLower(t.Name()) + "s" // exemplo: Pessoa -> pessoas
}

// acesso à coleção
func getCollection(doc interface{}) *mongo.Collection {
    collectionName := getCollectionName(doc)
    return config.Client.Database("orm_example").Collection(collectionName)
}

// 🔸 CREATE: Insere um documento no MongoDB
func Insert(document interface{}) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := getCollection(document)
    return collection.InsertOne(ctx, document)
}

// 🔸 RETRIEVE: Busca um único documento por filtro
func FindOne(docType interface{}, filter interface{}, result interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := getCollection(docType)
    err := collection.FindOne(ctx, filter).Decode(result)
    return err
}

// 🔸 UPDATE: Atualiza campos de um documento com base no filtro
func UpdateOne(docType interface{}, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := getCollection(docType)

    // usamos $set para definir os novos valores
    updateData := bson.M{"$set": update}

    return collection.UpdateOne(ctx, filter, updateData)
}

// 🔸 DELETE: Remove um documento com base no filtro
func DeleteOne(docType interface{}, filter interface{}) (*mongo.DeleteResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := getCollection(docType)
    return collection.DeleteOne(ctx, filter)
}
