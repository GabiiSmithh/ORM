package handle

import (
	"bufio"
	"fmt"
	"go-mongo-orm/models"
	"go-mongo-orm/orm"
	"strconv"
	"strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/google/uuid"
)

// Inserção transacional de Livro
func InsertLivro(livro models.Livro) func(sessCtx mongo.SessionContext) error {
	return func(sessCtx mongo.SessionContext) error {
		_, err := orm.GetCollection(livro).InsertOne(sessCtx, livro)
		return err
	}
}

// Atualização transacional de Livro
func UpdateLivro(id string, updateData map[string]interface{}) func(sessCtx mongo.SessionContext) error {
	return func(sessCtx mongo.SessionContext) error {
		filter := bson.M{"_id": id}
		update := bson.M{"$set": updateData}
		_, err := orm.GetCollection(models.Livro{}).UpdateOne(sessCtx, filter, update)
		return err
	}
}

// Remoção transacional de Livro
func DeleteLivro(id string) func(sessCtx mongo.SessionContext) error {
	return func(sessCtx mongo.SessionContext) error {
		filter := bson.M{"_id": id}
		_, err := orm.GetCollection(models.Livro{}).DeleteOne(sessCtx, filter)
		return err
	}
}

func HandleLivro(op string, reader *bufio.Reader) {
	switch op {
	case "1":
		fmt.Print("ISBN: ")
		isbn, _ := reader.ReadString('\n')
		isbn = strings.TrimSpace(isbn)

		fmt.Print("Título: ")
		titulo, _ := reader.ReadString('\n')
		titulo = strings.TrimSpace(titulo)

		fmt.Print("Autor: ")
		autor, _ := reader.ReadString('\n')
		autor = strings.TrimSpace(autor)

		fmt.Print("Ano de publicação: ")
		anoStr, _ := reader.ReadString('\n')
		anoStr = strings.TrimSpace(anoStr)
		anoPublic, _ := strconv.Atoi(anoStr)

		id := uuid.New().String()

		livro := models.Livro{
			ID:        id,
			ISBN:      isbn,
			Titulo:    titulo,
			Autor:     autor,
			AnoPublic: anoPublic,
		}

		res, err := orm.Insert(livro)
		if err != nil {
			fmt.Println("Erro ao inserir:", err)
		} else {
			fmt.Println("Livro inserido com ID:", res.InsertedID)
		}

	case "2":
		fmt.Print("Digite o nome para buscar: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		opts := orm.QueryOptions{
			Filter: map[string]interface{}{"titulo": nome},
		}
		var livros []models.Livro
		err := orm.FindCustom(models.Livro{}, opts, &livros)

		if err != nil {
			fmt.Println("Erro ao buscar:", err)
		} else {
			for _, p := range livros {
				fmt.Printf("Encontrado: %+v\n", p)
			}
		}

	case "3":
		fmt.Print("Digite o ISBN do livro para atualizar: ")
		isbn, _ := reader.ReadString('\n')
		isbn = strings.TrimSpace(isbn)

		fmt.Print("Novo título: ")
		titulo, _ := reader.ReadString('\n')
		titulo = strings.TrimSpace(titulo)

		filter := map[string]interface{}{"isbn": isbn}
		update := map[string]interface{}{"titulo": titulo}

		res, err := orm.UpdateOne(models.Livro{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4":
		fmt.Print("Digite o ISBN do livro para deletar: ")
		isbn, _ := reader.ReadString('\n')
		isbn = strings.TrimSpace(isbn)

		filter := map[string]interface{}{"isbn": isbn}
		res, err := orm.DeleteOne(models.Livro{}, filter)
		if err != nil {
			fmt.Println("Erro ao deletar:", err)
		} else {
			fmt.Printf("Deletados: %d\n", res.DeletedCount)
		}

	default:
		fmt.Println("Operação inválida")
	}
}

