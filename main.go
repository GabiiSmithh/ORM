package main

import (
	"bufio"
	"fmt"
	"go-mongo-orm/config"
	"go-mongo-orm/orm"
	"go-mongo-orm/utils/handle"
	"os"
	"strings"
)

func main() {
    config.Connect()

    // Chama a criação dos índices (uma vez só, no início)
	orm.EnsureIndexes()

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
    handle.HandlePessoa(opChoice, reader)
case "2":
    handle.HandleProduto(opChoice, reader)
case "3":
    handle.HandleLivro(opChoice, reader)
default:
    fmt.Println("Opção de modelo inválida")
}
