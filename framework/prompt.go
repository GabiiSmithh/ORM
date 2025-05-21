package framework

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
		parts := strings.Split(field, ",")
		fieldName := strings.TrimSpace(parts[0])
		fieldType := strings.TrimSpace(parts[1])
		val := PromptString(fmt.Sprintf("Valor para '%s' (%s): ", fieldName, fieldType))

		switch fieldType {
		case "int":
			parsed, _ := strconv.Atoi(val)
			doc[fieldName] = parsed
		case "float":
			parsed, _ := strconv.ParseFloat(val, 64)
			doc[fieldName] = parsed
		case "date":
			parsed, err := time.Parse("02/01/2006", val)
			if err != nil {
				fmt.Println("Formato de data inválido. Use dd/mm/aaaa.")
				continue
			}
			doc[fieldName] = parsed

		default:
			doc[fieldName] = val
		}
	}
	return doc
}

// Converte uma string para o tipo adequado (string, int, float, date)
func ConvertValue(input string, tipo string) (interface{}, error) {
	switch tipo {
	case "string":
		return input, nil
	case "int":
		return strconv.Atoi(input)
	case "float":
		return strconv.ParseFloat(input, 64)
	case "date":
		// Esperado: formato dd/mm/aaaa
		layout := "02/01/2006"
		t, err := time.Parse(layout, input)
		if err != nil {
			return nil, errors.New("data inválida, use o formato dd/mm/aaaa")
		}
		// Remove parte da hora mantendo apenas a data
		onlyDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		return onlyDate, nil
	default:
		return nil, errors.New("tipo desconhecido: " + tipo)
	}
}
