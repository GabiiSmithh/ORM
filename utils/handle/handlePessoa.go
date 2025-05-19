package handle

import (
	"bufio"
	"fmt"
	"go-mongo-orm/models"
	"go-mongo-orm/orm"
	"strconv"
	"strings"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlePessoa(op string, reader *bufio.Reader) {
	switch op {
	case "1": // Criar
		fmt.Print("Nome: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		fmt.Print("Data de nascimento (dd/mm/aaaa): ")
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)

		fmt.Print("CPF: ")
		cpf, _ := reader.ReadString('\n')
		cpf = strings.TrimSpace(cpf)

		fmt.Print("Telefone: ")
		tel, _ := reader.ReadString('\n')
		tel = strings.TrimSpace(tel)

		fmt.Print("Altura em cm (ex: 175): ")
		altStr, _ := reader.ReadString('\n')
		altStr = strings.TrimSpace(altStr)
		alt, _ := strconv.ParseFloat(altStr, 64)

		id := uuid.New().String()

		p := models.Pessoa{
			ID:       id,
			Nome:     nome,
			DataNasc: data,
			CPF:      cpf,
			Telefone: tel,
			Altura:   alt,
		}
		res, err := orm.Insert(p)
		if err != nil {
			fmt.Println("Erro ao inserir:", err)
		} else {
			fmt.Println("Pessoa inserida com ID:", res.InsertedID)
		}

	case "2": // Buscar
		fmt.Print("Digite o nome para buscar: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		opts := orm.QueryOptions{
			Filter: map[string]interface{}{"nome": nome},
		}

		var pessoas []models.Pessoa
		err := orm.FindCustom(models.Pessoa{}, opts, &pessoas)
		if err != nil {
			fmt.Println("Erro ao buscar:", err)
		} else {
			for _, p := range pessoas {
				fmt.Printf("Encontrado: %+v\n", p)
			}
		}

	case "3": // Atualizar
		fmt.Print("Digite o CPF da pessoa para atualizar: ")
	 	cpf, _ := reader.ReadString('\n')
	 	cpf = strings.TrimSpace (cpf)

		fmt.Print("Novo telefone: ")
		tel, _ := reader.ReadString('\n')
		tel = strings.TrimSpace(tel)

		filter := map[string]interface{}{"cpf": cpf}
		update := map[string]interface{}{"telefone": tel}

		res, err := orm.UpdateOne(models.Pessoa{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4": // Deletar
		fmt.Print("Digite o CPF da pessoa para deletar: ")
	 	cpf, _ := reader.ReadString('\n')
	 	cpf = strings.TrimSpace (cpf)

		filter := map[string]interface{}{"cpf": cpf}
		res, err := orm.DeleteOne(models.Pessoa{}, filter)
		if err != nil {
			fmt.Println("Erro ao deletar:", err)
		} else {
			fmt.Printf("Deletados: %d\n", res.DeletedCount)
		}

	default:
		fmt.Println("Operação inválida")
	}
}

// InserirPessoa insere uma nova pessoa no banco de dados
func InserirPessoa(sessCtx mongo.SessionContext, client *mongo.Client, nome string) error {
	collection := client.Database("orm_example").Collection("pessoas")
	_, err := collection.InsertOne(sessCtx, bson.M{"nome": nome})
	return err
}

// AtualizarPessoa atualiza o nome de uma pessoa
func AtualizarPessoa(sessCtx mongo.SessionContext, client *mongo.Client, id string, nome string) error {
	collection := client.Database("orm_example").Collection("pessoas")
	_, err := collection.UpdateOne(sessCtx, bson.M{"_id": id}, bson.M{"$set": bson.M{"nome": nome}})
	return err
}

// DeletarPessoa deleta uma pessoa pelo ID
func DeletarPessoa(sessCtx mongo.SessionContext, client *mongo.Client, id string) error {
	collection := client.Database("orm_example").Collection("pessoas")
	_, err := collection.DeleteOne(sessCtx, bson.M{"_id": id})
	return err
}