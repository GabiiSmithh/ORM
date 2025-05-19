package main

import (
	"bufio"
	"fmt"
	"go-mongo-orm/config"
	"go-mongo-orm/orm"
	"go-mongo-orm/utils/handle"
    "go.mongodb.org/mongo-driver/mongo"
	"os"
	"strings"
	"github.com/google/uuid"
	"go-mongo-orm/models"

)



func executarTransacaoExemplo() {
	client := orm.GetMongoClient() // Sua função para obter o client do MongoDB

	// Inicializa o leitor para capturar entradas do usuário
	reader := bufio.NewReader(os.Stdin)

	var ops []orm.TxOperation

	for {
		// Menu de operações
		fmt.Println("Escolha uma operação:")
		fmt.Println("1 - Inserir Pessoa")
		fmt.Println("2 - Atualizar Pessoa")
		fmt.Println("3 - Deletar Pessoa")
		fmt.Println("4 - Inserir Produto")
		fmt.Println("5 - Atualizar Produto")
		fmt.Println("6 - Deletar Produto")
		fmt.Println("7 - Sair")
		fmt.Println("8 - Inserir Livro")
		fmt.Println("9 - Atualizar Livro")
		fmt.Println("10 - Deletar Livro")
		fmt.Print("Escolha a opção: ")

		// Leitura da escolha do usuário
		opcao, _ := reader.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		// Processa a opção escolhida
		switch opcao {
		case "1": // Inserir Pessoa
			fmt.Print("Nome da pessoa: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.InserirPessoa(sessCtx, client, nome)
			})

		case "2": // Atualizar Pessoa
			fmt.Print("ID da pessoa: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			fmt.Print("Novo nome da pessoa: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.AtualizarPessoa(sessCtx, client, id, nome)
			})

		case "3": // Deletar Pessoa
			fmt.Print("ID da pessoa a ser deletada: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.DeletarPessoa(sessCtx, client, id)
			})

		case "4": // Inserir Produto
			fmt.Print("Nome do produto: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)
			fmt.Print("Preço do produto: ")
			var preco float64
			fmt.Scanf("%f\n", &preco)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.InserirProduto(sessCtx, client, nome, preco)
			})

		case "5": // Atualizar Produto
			fmt.Print("ID do produto: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			fmt.Print("Novo nome do produto: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)
			fmt.Print("Novo preço do produto: ")
			var preco float64
			fmt.Scanf("%f\n", &preco)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.AtualizarProduto(sessCtx, client, id, nome, preco)
			})

		case "6": // Deletar Produto
			fmt.Print("ID do produto a ser deletado: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			ops = append(ops, func(sessCtx mongo.SessionContext) error {
				return handle.DeletarProduto(sessCtx, client, id)
			})

		case "7": // Sair
			fmt.Println("Saindo...")
			return

		case "8": // Inserir Livro
			fmt.Print("ISBN do livro: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)

			fmt.Print("Título do livro: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Autor do livro: ")
			autor, _ := reader.ReadString('\n')
			autor = strings.TrimSpace(autor)

			fmt.Print("Ano de publicação: ")
			var ano int
			fmt.Scanf("%d\n", &ano)

			id := uuid.New().String()
			livro := models.Livro{
				ID:        id,
				ISBN:      isbn,
				Titulo:    titulo,
				Autor:     autor,
				AnoPublic: ano,
			}

			ops = append(ops, handle.InsertLivro(livro))

		case "9": // Atualizar Livro
			fmt.Print("ID do livro: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Novo título: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			updateData := map[string]interface{}{
				"titulo": titulo,
			}

			ops = append(ops, handle.UpdateLivro(id, updateData))

		case "10": // Deletar Livro
			fmt.Print("ID do livro a ser deletado: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			ops = append(ops, handle.DeleteLivro(id))

		default:
			fmt.Println("Opção inválida! Tente novamente")
			continue
		}

		// Pergunta se o usuário deseja realizar mais operações ou executar a transação
		fmt.Print("Deseja realizar mais operações? (s/n): ")
		continuar, _ := reader.ReadString('\n')
		if strings.TrimSpace(continuar) != "s" {
			break
		}
	}

	// Executa a transação com todas as operações selecionadas
	err := orm.RunTransaction(client, ops)
	if err != nil {
		fmt.Println("Erro na transação:", err)
	} else {
		fmt.Println("Transação realizada com sucesso.")
	}
}


func main() {
    executarTransacaoExemplo()
    config.Connect()

    // Chama a criação dos índices (uma vez só, no início)
	orm.EnsureIndexes()

    reader := bufio.NewReader(os.Stdin)

	handle.RodarTestesLivro()

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
			handle.HandlePessoa(opChoice, reader)
		case "2":
			handle.HandleProduto(opChoice, reader)
		case "3":
			handle.HandleLivro(opChoice, reader)
		default:
			fmt.Println("Opção de modelo inválida")
		}


	}

}

