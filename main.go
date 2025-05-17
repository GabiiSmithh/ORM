package main

import (
	"bufio"
	"fmt"
	"go-mongo-orm/config"
	"go-mongo-orm/models"
	"go-mongo-orm/orm"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	config.Connect()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nEscolha o modelo:")
		fmt.Println("1 - Pessoa")
		fmt.Println("2 - Produto")
		fmt.Println("3 - Livro")
		fmt.Println("0 - Sair")
		fmt.Print("Opção: ")
		modelChoice, _ := reader.ReadString('\n')
		modelChoice = strings.TrimSpace(modelChoice)

		if modelChoice == "0" {
			fmt.Println("Saindo...")
			break
		}

		fmt.Println("\nEscolha a operação:")
		fmt.Println("1 - Criar")
		fmt.Println("2 - Buscar (Read)")
		fmt.Println("3 - Atualizar")
		fmt.Println("4 - Deletar")
		fmt.Print("Opção: ")
		opChoice, _ := reader.ReadString('\n')
		opChoice = strings.TrimSpace(opChoice)

		switch modelChoice {
		case "1":
			handlePessoa(opChoice, reader)
		case "2":
			handleProduto(opChoice, reader)
		case "3":
			handleLivro(opChoice, reader)
		default:
			fmt.Println("Opção de modelo inválida")
		}
	}
}

// --- Pessoa ---
func handlePessoa(op string, reader *bufio.Reader) {
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

		filter := map[string]interface{}{"nome": nome}
		var pessoas []models.Pessoa
		err := orm.FindMany(models.Pessoa{}, filter, &pessoas)
		if err != nil {
			fmt.Println("Erro ao buscar:", err)
		} else {
			for _, p := range pessoas {
				fmt.Printf("Encontrado: %+v\n", p)
			}
		}

	case "3": // Atualizar
		fmt.Print("Digite o ID da pessoa para atualizar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		fmt.Print("Novo telefone: ")
		tel, _ := reader.ReadString('\n')
		tel = strings.TrimSpace(tel)

		filter := map[string]interface{}{"_id": id}
		update := map[string]interface{}{"telefone": tel}

		res, err := orm.UpdateOne(models.Pessoa{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4": // Deletar
		fmt.Print("Digite o ID da pessoa para deletar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		filter := map[string]interface{}{"_id": id}
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

// --- Produto ---
func handleProduto(op string, reader *bufio.Reader) {
	switch op {
	case "1":
		fmt.Print("Código do produto: ")
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
		fmt.Print("Digite o ID do produto para atualizar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		fmt.Print("Nova quantidade: ")
		qtdStr, _ := reader.ReadString('\n')
		qtdStr = strings.TrimSpace(qtdStr)
		qtd, _ := strconv.Atoi(qtdStr)

		filter := map[string]interface{}{"_id": id}
		update := map[string]interface{}{"quantidade": qtd}

		res, err := orm.UpdateOne(models.Produto{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4":
		fmt.Print("Digite o ID do produto para deletar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		filter := map[string]interface{}{"_id": id}
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

// --- Livro ---
func handleLivro(op string, reader *bufio.Reader) {
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

		filter := map[string]interface{}{"nome": nome}
		var livros []models.Livro
		err := orm.FindMany(models.Livro{}, filter, &livros)
		if err != nil {
			fmt.Println("Erro ao buscar:", err)
		} else {
			for _, p := range livros {
				fmt.Printf("Encontrado: %+v\n", p)
			}
		}

	case "3":
		fmt.Print("Digite o ID do livro para atualizar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		fmt.Print("Novo título: ")
		titulo, _ := reader.ReadString('\n')
		titulo = strings.TrimSpace(titulo)

		filter := map[string]interface{}{"_id": id}
		update := map[string]interface{}{"titulo": titulo}

		res, err := orm.UpdateOne(models.Livro{}, filter, update)
		if err != nil {
			fmt.Println("Erro ao atualizar:", err)
		} else {
			fmt.Printf("Modificados: %d\n", res.ModifiedCount)
		}

	case "4":
		fmt.Print("Digite o ID do livro para deletar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		filter := map[string]interface{}{"_id": id}
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
