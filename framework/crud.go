package framework

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Colecao struct { // Colecao representa uma coleção no MongoDB -> Genérica para qualquer coleção
	Colecao *mongo.Collection
}

func NovaColecao(db *mongo.Database, nomeColecao string) *Colecao { // Cria a coleção a partir do banco de dados, retorna uma instância de Colecao
	return &Colecao{
		Colecao: db.Collection(nomeColecao),
	}
}

// CREATE
func (r *Colecao) Inserir(dado map[string]interface{}) (*mongo.InsertOneResult, error) { // Insere um documento a partir um mapeamento de modelo genérico
	return r.Colecao.InsertOne(context.Background(), dado)
}

// BuscarComFiltroOrdenacao permite buscar documentos com filtro e ordenação personalizados
func (r *Colecao) BuscarComFiltroOrdenacao(filtro map[string]interface{}, ordenacao map[string]int) ([]bson.M, error) {
	// Transforma os mapas recebidos para tipos BSON compatíveis com o driver do MongoDB
	filter := bson.M(filtro)
	sort := bson.D{}
	for campo, ordem := range ordenacao {
		sort = append(sort, bson.E{Key: campo, Value: ordem})
	}

	// Define as opções da consulta com ordenação
	opcoes := &options.FindOptions{
		Sort: sort,
	}

	// Executa a consulta com filtro e opções
	cursor, err := r.Colecao.Find(context.Background(), filter, opcoes)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resultados []bson.M
	for cursor.Next(context.Background()) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		resultados = append(resultados, doc)
	}

	return resultados, nil
}


// RETRIEBE (ONE)
func (r *Colecao) BuscarPorID(id primitive.ObjectID) (bson.M, error) { // Busca um documento específico a partir do ID
	var resultado bson.M
	err := r.Colecao.FindOne(context.Background(), bson.M{"_id": id}).Decode(&resultado)
	return resultado, err
}

// UPDATE
func (r *Colecao) Atualizar(id primitive.ObjectID, atualizacoes map[string]interface{}) (*mongo.UpdateResult, error) { // Atualiza um documento de acordo com um id
	return r.Colecao.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": atualizacoes}, // Atualiza o documento com os dados passados
	)
}

// DELETE
func (r *Colecao) Remover(id primitive.ObjectID) (*mongo.DeleteResult, error) { // Remove um documento a partir do ID
	return r.Colecao.DeleteOne(context.Background(), bson.M{"_id": id})
}
