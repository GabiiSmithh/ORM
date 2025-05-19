package handle

import (
	"fmt"
	"go-mongo-orm/models"
	"go-mongo-orm/orm"

	"github.com/google/uuid"
)

// Teste 1 – Inserção com tentativa de duplicidade
func ExemploInsercaoComDuplicidade() {
	fmt.Println("Teste 1 – Inserção com tentativa de duplicidade:")

	livro1 := models.Livro{
		ID:        uuid.New().String(),
		ISBN:      "111-111",
		Titulo:    "Livro Original",
		Autor:     "Autor A",
		AnoPublic: 2020,
	}

	livro2 := models.Livro{
		ID:        uuid.New().String(),
		ISBN:      "111-111", // mesmo ISBN
		Titulo:    "Livro Duplicado",
		Autor:     "Autor B",
		AnoPublic: 2021,
	}

	if _, err := orm.Insert(livro1); err != nil {
		fmt.Println("Erro ao inserir o livro1:", err)
	} else {
		fmt.Println("Livro1 inserido com sucesso")
	}

	if _, err := orm.Insert(livro2); err != nil {
		fmt.Println("Erro ao inserir livro duplicado (esperado):", err)
	} else {
		fmt.Println("Livro2 inserido (isso não deveria acontecer!)")
	}
}

// Teste 2 – Atualização de título pelo ISBN
func ExemploAtualizacaoTitulo() {
	fmt.Println("Teste 2 – Atualização de título pelo ISBN:")

	filter := map[string]interface{}{
		"isbn": "111-111",
	}
	update := map[string]interface{}{
		"titulo": "Livro Atualizado com Sucesso",
	}

	res, err := orm.UpdateOne(models.Livro{}, filter, update)
	if err != nil {
		fmt.Println("Erro ao atualizar:", err)
	} else {
		fmt.Printf("Títulos atualizados: %d\n", res.ModifiedCount)
	}
}

// Teste 3 – Deleção de um livro pelo ISBN
func ExemploDelecaoLivro() {
	fmt.Println("Teste 3 – Deleção de um livro:")

	filter := map[string]interface{}{
		"isbn": "111-111",
	}

	res, err := orm.DeleteOne(models.Livro{}, filter)
	if err != nil {
		fmt.Println("Erro ao deletar:", err)
	} else {
		fmt.Printf("Livros deletados: %d\n", res.DeletedCount)
	}
}

// Função para rodar todos os testes
func RodarTestesLivro() {
	fmt.Println("==== INICIANDO TESTES COM LIVROS ====")
	ExemploInsercaoComDuplicidade()
	fmt.Println()
	ExemploAtualizacaoTitulo()
	fmt.Println()
	ExemploDelecaoLivro()
	fmt.Println("==== FIM DOS TESTES ====")
}