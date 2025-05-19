package handle

import (
	"bufio"
	"fmt"
	"go-mongo-orm/models"
	"go-mongo-orm/orm"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func HandleProduto(op string, reader *bufio.Reader) {
	switch op {
	case "1":
		fmt.Print("Código do produto: ")
		codigo, _ := reader.ReadString('\n')
		codigo = strings.TrimSpace(codigo)

		fmt.Print("Categoria do produto: ")
		categoria, _ := reader.ReadString('\n')
		categoria = strings.TrimSpace(categoria)

		fmt.Print("Nome do produto: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		fmt.Print("Preço: ")
		precoStr, _ := reader.ReadString('\n')
		precoStr = strings.TrimSpace(precoStr)
		preco, _ := strconv.ParseFloat(precoStr, 64)

		fmt.Print("Quantidade: ")
		qtdStr, _ := reader.ReadString('\n')
		qtdStr = strings.TrimSpace(qtdStr)
		qtd, _ := strconv.Atoi(qtdStr)

		id := uuid.New().String()

		p := models.Produto{
			ID:         id,
			Codigo:     codigo,
			Nome:       nome,
			Categoria:  categoria,
			Preco:      preco,
			Quantidade: qtd,
		}
		res, err := orm.Insert(p)
		if err != nil {
			fmt.Println("Erro ao inserir:", err)
		} else {
			fmt.Println("Produto inserido com ID:", res.InsertedID)
		}

	case "2":
		fmt.Print("Digite o nome para buscar: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		filter := map[string]interface{}{"nome": nome}
		var produtos []models.Produto
		err := orm.FindMany(models.Produto{}, filter, &produtos)
		if err != nil {
			fmt.Println("Erro ao buscar:", err)
		} else {
			for _, p := range produtos {
				fmt.Printf("Encontrado: %+v\n", p)
			}
		}

	case "3":
		fmt.Print("Digite o código do produto para atualizar: ")
		codigo, _ := reader.ReadString('\n')
		codigo = strings.TrimSpace(codigo)

		fmt.Print("Nova quantidade: ")
		qtdStr, _ := reader.ReadString('\n')
		qtdStr = strings.TrimSpace(qtdStr)
		qtd, _ := strconv.Atoi(qtdStr)

		filter := map[string]interface{}{"codigo": codigo}
		update := map[string]interface{}{"quantidade": qtd}

		res, err := orm.UpdateOne(models.Produto{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4":
		fmt.Print("Digite o ID do produto para deletar: ")
		codigo, _ := reader.ReadString('\n')
		codigo = strings.TrimSpace(codigo)

		filter := map[string]interface{}{"codigo": codigo}
		res, err := orm.DeleteOne(models.Produto{}, filter)
		if err != nil {
			fmt.Println("Erro ao deletar:", err)
		} else {
			fmt.Printf("Deletados: %d\n", res.DeletedCount)
		}

	default:
		fmt.Println("Operação inválida")
	}
}