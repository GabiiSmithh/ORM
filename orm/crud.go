package orm

import (
	"context"
	"fmt"
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
func FindMany(model interface{}, filter interface{}, result interface{}) error {
	coll := getCollection(model)

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	slice := reflect.ValueOf(result).Elem()
	elemType := slice.Type().Elem()

	for cursor.Next(context.TODO()) {
		elemPtr := reflect.New(elemType)
		err := cursor.Decode(elemPtr.Interface())
		if err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, elemPtr.Elem()))
	}

	return cursor.Err()
}

// 🔸 UPDATE: Atualiza campos de um documento com base no filtro
func UpdateOne(docType interface{}, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection(docType)
	updateData := bson.M{"$set": update}

	// Debug temporário
	fmt.Println("⛏️ Filtro:", filter)
	fmt.Println("🛠️ Update:", updateData)

	return collection.UpdateOne(ctx, filter, updateData)
}

// 🔸 DELETE: Remove um documento com base no filtro
func DeleteOne(docType interface{}, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getCollection(docType)
	return collection.DeleteOne(ctx, filter)
}
