package framework

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptString(prompt string) string { // Função para solicitar informações ao usuário
	reader := bufio.NewReader(os.Stdin) // Cria um leitor ligado à entrada padrão (teclado)
	fmt.Print(prompt)                   // Exibe o texto do prompt para o usuário

	input, _ := reader.ReadString('\n') // Lê até o usuário apertar ENTER
	return strings.TrimSpace(input)     // Remove espaços e quebras de linha do final
}

func PromptDocumentFields(fields []string) map[string]interface{} {
	doc := make(map[string]interface{})
	for _, field := range fields {
		val := PromptString(fmt.Sprintf("Valor para %s: ", field))
		doc[field] = val
	}
	return doc
}


