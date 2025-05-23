package main

import (
	"context"
	"fmt"
	"log"

	"go-mongo-orm/config"
	"go-mongo-orm/framework"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Configurações do MongoDB
	uri := "mongodb://localhost:27017"
	nomeDB := "meudb"
	nomeColecao := "registros"

	// Conectando ao MongoDB
	cliente, err := config.Conectar(uri)
	if err != nil {
		log.Fatal("Erro na conexão com o banco:", err)
	}
	defer cliente.Disconnect(context.TODO())

	// Acessando a coleção
	db := cliente.Database(nomeDB)
	repo := framework.NovaColecao(db, nomeColecao)


	// Exemplo de inserção de um documento
	registro := map[string]interface{}{
		"nome":  "Gabriela",
		"curso": "Ciencia da Computação",
		"idade": 22,
	}
	resultado, err := repo.Inserir(registro)
	if err != nil {
		log.Fatal("Erro ao inserir:", err)
	}
	fmt.Println("ID do documento inserido:", resultado.InsertedID)

	// Exemplo de atualização de um documento recém-inserido
	id := resultado.InsertedID.(primitive.ObjectID)
	_, err = repo.Atualizar(id, map[string]interface{}{"curso": "Engenharia"})
	if err != nil {
		log.Fatal("Erro ao atualizar:", err)
	}

	// Exemplo de busca com filtro e ordenação
	filtro := map[string]interface{}{
		"curso": "Engenharia",
	}
	ordenacao := map[string]int{
		"idade": -1, // ordem decrescente
	}

	documentosFiltrados, err := repo.BuscarComFiltroOrdenacao(filtro, ordenacao)
	if err != nil {
		log.Fatal("Erro ao buscar com filtro e ordenação:", err)
	}
	fmt.Println("Documentos encontrados com filtro e ordenação:", documentosFiltrados)
}
