package main

import (
	"context"
	"fmt"  
	"log"   
	"time" 

	"go-mongo-orm/config"   
	"go-mongo-orm/framework"
	"go-mongo-orm/models" 

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	uri := "mongodb://localhost:27017"
	nomeDB := "meudb"
	nomeColecao := "produtos"

	fmt.Println("Iniciando...") // Testando se está tentando conectar

	// Tenta estabelecer conexão com o MongoDB usando o pacote config
	cliente, err := config.Conectar(uri)
	if err != nil {
		log.Fatal("Erro na conexão com o banco:", err)
	}

	defer cliente.Disconnect(context.TODO())	// Garante que a conexão será fechada no final da função main

	// Seleciona o banco de dados com o nome especificado
	db := cliente.Database(nomeDB)
	// Cria o repositório para manipular a coleção de produtos
	repo := framework.NovaColecao(db, nomeColecao)

	// Cria uma nova instância do produto com dados para inserção
	produto := models.Produto{
		Nome:        "Teclado Mecânico",                   // Nome do produto
		Categoria:   "Periféricos",                         // Categoria do produto
		Preco:       250.75,                                // Preço do produto
		Disponivel:  true,                                  // Disponibilidade no estoque
		DataEntrada: primitive.NewDateTimeFromTime(time.Now()), // Data de entrada atual no formato BSON
	}

	// Insere o produto criado na coleção MongoDB usando o método do framework
	resultado, err := repo.Inserir(produto)
	if err != nil {
		log.Fatal("Erro ao inserir produto:", err)
	}
	fmt.Println("ID do produto inserido:", resultado.InsertedID)

	// Converte o ID inserido para o tipo ObjectID do MongoDB para manipular depois
	id := resultado.InsertedID.(primitive.ObjectID)
	atualizacao := map[string]interface{}{ 	// Define um mapa com a atualização desejada (novo preço)
		"preco": 230.50, // Novo valor do preço para atualizar
	}

	// Executa a atualização do produto no banco, buscando pelo ID
	_, err = repo.Atualizar(id, atualizacao)
	if err != nil {
		log.Fatal("Erro ao atualizar produto:", err)
	}

	// Define um filtro para buscar somente produtos que estão disponíveis e na categoria "Periféricos"
	filtro := map[string]interface{}{
		"categoria":  "Periféricos",
		"disponivel": true,
	}
	// Define a ordenação dos resultados pelo campo preço em ordem decrescente (-1)
	ordenacao := map[string]int{
		"preco": -1,
	}

	// Busca os produtos na coleção usando filtro e ordenação definidos
	produtosFiltrados, err := repo.BuscarComFiltroOrdenacao(filtro, ordenacao)
	if err != nil {
		// Se der erro na busca, para a execução e mostra o erro
		log.Fatal("Erro ao buscar produtos:", err)
	}
	// Imprime no console os produtos encontrados
	fmt.Println("Produtos encontrados:", produtosFiltrados)
}
