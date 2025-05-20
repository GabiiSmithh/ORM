package framework

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"

	"go-mongo-orm/config"
)

// Create
func InsertDocument(collectionName string, data map[string]interface{}) error {
	coll := config.GetCollection(collectionName)   // Obtém a coleção pelo nome
	_, err := coll.InsertOne(context.TODO(), data) // Insere o documento
	return err                                     // Caso haja erro
}

// Retrieve
func FindDocuments(collectionName string, filter bson.M) ([]bson.M, error) {
	coll := config.GetCollection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter) // executa a busca com filtro
	if err != nil {
		return nil, err
	}

	var results []bson.M // slice para armazenar os documentos encontrados
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func UpdateDocumentByPrimaryKey(collectionName, primaryKey, keyValue string, update bson.M) error {
	coll := config.GetCollection(collectionName)

	filter := bson.M{primaryKey: keyValue} // usa string direto
	updateDoc := bson.M{"$set": update}

	_, err := coll.UpdateOne(context.TODO(), filter, updateDoc)
	return err
}

func DeleteDocumentByPrimaryKey(collectionName, primaryKey, keyValue string) error {
	coll := config.GetCollection(collectionName)

	filter := bson.M{primaryKey: keyValue} // usa string direto
	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}

// Salva o esquema de uma nova coleção
func SaveCollectionSchema(collectionName string, fields []string, primaryKey string) error {
	schemaColl := config.GetCollection("schemas")
	doc := bson.M{
		"collection":  collectionName,
		"fields":      fields,
		"primary_key": primaryKey,
	}

	// Remove schema anterior, se existir
	_, err := schemaColl.DeleteMany(context.TODO(), bson.M{"collection": collectionName})
	if err != nil {
		return err
	}
	// Insere o novo schema
	_, err = schemaColl.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	// Cria índice único no campo da chave primária
	coll := config.GetCollection(collectionName)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: primaryKey, Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return fmt.Errorf("erro ao criar índice único na chave primária '%s': %v", primaryKey, err)
	}

	return nil
}

// Carrega o esquema e chave primária de uma coleção
func GetCollectionSchema(collectionName string) ([]string, string, error) {
	schemaColl := config.GetCollection("schemas")
	var result struct {
		Fields     []string `bson:"fields"`
		PrimaryKey string   `bson:"primary_key"`
	}
	err := schemaColl.FindOne(context.TODO(), bson.M{"collection": collectionName}).Decode(&result)
	if err != nil {
		return nil, "", err
	}
	return result.Fields, result.PrimaryKey, nil
}
