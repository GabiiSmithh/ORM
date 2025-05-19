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
	"go.mongodb.org/mongo-driver/mongo/options"
)


// Exporta o client do MongoDB para uso externo
func GetMongoClient() *mongo.Client {
	return config.Client
}

// obtém o nome da coleção a partir do nome da struct
func getCollectionName(doc interface{}) string {
	t := reflect.TypeOf(doc) // obtém o tipo da struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return strings.ToLower(t.Name()) + "s" // gera o nome da coleção
}

// acesso à coleção
func GetCollection(doc interface{}) *mongo.Collection {
	collectionName := getCollectionName(doc) // obtém o nome da coleção
	return config.Client.Database("orm_example").Collection(collectionName) // acessa a coleção em determinado banco
}

// CREATE: Insere um documento
func Insert(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection(document) // Obtém a coleção correspondente ao tipo do documento
	return collection.InsertOne(ctx, document) // Realiza a inserção
}

// RETRIEVE: Busca um documento a partir de filtros, ordenação e seleção de campos
func FindCustom(model interface{}, opts QueryOptions, result interface{}) error {
	coll := GetCollection(model) // Obtém a coleção correspondente ao tipo do modelo

	QueryOptions := options.Find()
	if opts.Projection != nil{
		QueryOptions.SetProjection(opts.Projection)
	}
	if opts.Sort != nil {
		QueryOptions.SetSort(opts.Sort)
	}

	cursor, err := coll.Find(context.TODO(), opts.Filter, QueryOptions) // Realiza a busca com o filtro
	if err != nil { // Verifica se houve erro na busca
		return err
	}
	defer cursor.Close(context.TODO())

	slice := reflect.ValueOf(result).Elem() // Obtém o valor do slice passado como resultado
	elemType := slice.Type().Elem() // Obtém o tipo do elemento do slice

	for cursor.Next(context.TODO()) { // Itera sobre os resultados	
		elemPtr := reflect.New(elemType)
		err := cursor.Decode(elemPtr.Interface())
		if err != nil {
			return err
		}
		slice.Set(reflect.Append(slice, elemPtr.Elem())) // Adiciona o elemento decodificado ao slice
	}

	return cursor.Err()
}

// UPDATE: Atualiza campos de um documento com base no filtro
func UpdateOne(docType interface{}, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection(docType) // Obtém a coleção correspondente ao tipo do documento
	updateData := bson.M{"$set": update} // Cria o documento de atualização

	// Debug temporário
	fmt.Println("Filtro:", filter) // Exibe o filtro
	fmt.Println("Update:", updateData) // Exibe os dados a serem atualizados

	return collection.UpdateOne(ctx, filter, updateData) // Realiza a atualização
}

// DELETE: Remove um documento com base no filtro
func DeleteOne(docType interface{}, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := GetCollection(docType) // Obtém a coleção correspondente ao tipo do documento
	return collection.DeleteOne(ctx, filter) // Realiza a remoção
}
