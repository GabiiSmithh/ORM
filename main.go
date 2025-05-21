package main

import (
	"context"
	"fmt"
	"go-mongo-orm/config"
	"go-mongo-orm/framework"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// Prompta valores com base em campos definidos no schema
func PromptValuesFromSchema(fields []string) map[string]interface{} {
	doc := make(map[string]interface{})
	for _, field := range fields {
		val := framework.PromptString(fmt.Sprintf("Valor para '%s': ", field))
		doc[field] = val
	}
	return doc
}

func main() {
	mongoURI := "mongodb://localhost:27017"
	dbName := "orm-test"
	fmt.Println("Conectando ao MongoDB...")
	config.Connect(mongoURI, dbName)

	escolha := framework.PromptString("Criar nova coleção (c) ou usar existente (u)? ")
	var collectionName string

	if escolha == "c" {
		collectionName = framework.PromptString("Nome da nova coleção: ")
		fmt.Println("Informe os campos no formato: nomeCampo, tipoCampo (string, int, float, date). Deixe vazio para terminar.")
		var fields []string
		for {
			input := framework.PromptString("Campo: ")
			if input == "" {
				break
			}

			parts := strings.Split(strings.TrimSpace(input), ",")
			if len(parts) != 2 {
				fmt.Println("Formato inválido. Use: nomeCampo, tipoCampo")
				continue
			}

			field := strings.TrimSpace(parts[0])
			fieldType := strings.ToLower(strings.TrimSpace(parts[1]))

			if fieldType != "string" && fieldType != "int" && fieldType != "float" && fieldType != "date" {
				fmt.Println("Tipo inválido. Tipos permitidos: string, int, float, date")
				continue
			}

			fields = append(fields, field+","+fieldType)
		}

		primaryKey := framework.PromptString("Informe o nome do campo que será a chave primária: ")
		err := framework.SaveCollectionSchema(collectionName, fields, primaryKey)
		if err != nil {
			log.Fatal("Erro ao salvar esquema:", err)
		}
		fmt.Println("Esquema da coleção salvo com sucesso!")
	} else if escolha == "u" {
		// Listar coleções existentes
		schemaColl := config.GetCollection("schemas")
		cursor, err := schemaColl.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal("Erro ao buscar esquemas:", err)
		}
		var schemas []bson.M
		if err = cursor.All(context.TODO(), &schemas); err != nil {
			log.Fatal("Erro ao decodificar esquemas:", err)
		}
		if len(schemas) == 0 {
			fmt.Println("Nenhuma coleção encontrada. Crie uma nova coleção primeiro.")
			return
		}
		fmt.Println("Coleções disponíveis:")
		for _, schema := range schemas {
			fmt.Printf("- %s\n", schema["collection"])
		}
		collectionName = framework.PromptString("Nome da coleção existente: ")
		fields, primaryKey, err := framework.GetCollectionSchema(collectionName)
		if err != nil {
			log.Fatal("Erro ao obter esquema da coleção:", err)
		}

		// CRUD principal
		for {
			fmt.Println("\nEscolha a operação:")
			fmt.Println("1 - Inserir documento")
			fmt.Println("2 - Listar documentos")
			fmt.Println("3 - Atualizar documento por chave primária")
			fmt.Println("4 - Deletar documento por chave primária")
			fmt.Println("0 - Sair")

			opcao := framework.PromptString("Opção: ")

			switch opcao {
			case "1":
				doc := framework.PromptDocumentFields(fields)
				err := framework.InsertDocument(collectionName, doc)
				if err != nil {
					log.Println("Erro ao inserir documento:", err)
				} else {
					fmt.Println("Documento inserido com sucesso!")
				}
			case "2":
				docs, err := framework.FindDocuments(collectionName, bson.M{})
				if err != nil {
					log.Println("Erro ao buscar documentos:", err)
					continue
				}
				fmt.Printf("Documentos na coleção '%s':\n", collectionName)
				for i, d := range docs {
					fmt.Printf("%d) %s=%v, dados=%+v\n", i+1, primaryKey, d[primaryKey], d)
				}
			case "3":
				keyValue := framework.PromptString(fmt.Sprintf("Informe o valor da chave primária (%s) do documento para atualizar: ", primaryKey))

				// Extrai apenas os nomes dos campos do schema
				fieldNames := make([]string, 0)
				fieldTypes := make(map[string]string)
				for _, f := range fields {
					parts := strings.Split(strings.TrimSpace(f), ",")
					if len(parts) == 2 {
						name := strings.TrimSpace(parts[0])
						typ := strings.TrimSpace(parts[1])
						fieldNames = append(fieldNames, name)
						fieldTypes[name] = typ
					}
				}

				fmt.Println("Campos disponíveis para atualização:")
				for _, field := range fieldNames {
					fmt.Printf("- %s\n", field)
				}

				updateData := make(map[string]interface{})
				for {
					campo := framework.PromptString("Digite o nome do campo a ser alterado (ou pressione Enter para finalizar): ")
					if campo == "" {
						break
					}

					if _, ok := fieldTypes[campo]; !ok {
						fmt.Println("Campo inválido. Tente novamente.")
						continue
					}

					valorStr := framework.PromptString(fmt.Sprintf("Novo valor para '%s': ", campo))
					// Converte o valor com base no tipo definido
					valorConvertido, err := framework.ConvertValue(valorStr, fieldTypes[campo])
					if err != nil {
						fmt.Println("Erro na conversão do valor:", err)
						continue
					}
					updateData[campo] = valorConvertido
				}

				if len(updateData) == 0 {
					fmt.Println("Nenhum campo selecionado para atualização.")
				} else {
					err := framework.UpdateDocumentByPrimaryKey(collectionName, primaryKey, keyValue, updateData)
					if err != nil {
						log.Println("Erro ao atualizar documento:", err)
					} else {
						fmt.Println("Documento atualizado com sucesso!")
					}
				}

			case "4":
				keyValue := framework.PromptString(fmt.Sprintf("Informe o valor da chave primária (%s) do documento para deletar: ", primaryKey))
				err := framework.DeleteDocumentByPrimaryKey(collectionName, primaryKey, keyValue)
				if err != nil {
					log.Println("Erro ao deletar documento:", err)
				} else {
					fmt.Println("Documento deletado com sucesso!")
				}
			case "0":
				fmt.Println("Saindo...")
				return
			default:
				fmt.Println("Opção inválida!")
			}
		}

	}
}
